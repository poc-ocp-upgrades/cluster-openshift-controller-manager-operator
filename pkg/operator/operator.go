package operator

import (
	"fmt"
	"time"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	operatorapiv1 "github.com/openshift/api/operator/v1"
	operatorclientv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	operatorinformersv1 "github.com/openshift/client-go/operator/informers/externalversions/operator/v1"
	"github.com/openshift/cluster-openshift-controller-manager-operator/pkg/util"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
)

const (
	workQueueKey				= "key"
	workloadFailingCondition	= "WorkloadFailing"
)

type OpenShiftControllerManagerOperator struct {
	targetImagePullSpec		string
	operatorConfigClient	operatorclientv1.OperatorV1Interface
	kubeClient				kubernetes.Interface
	queue					workqueue.RateLimitingInterface
	rateLimiter				flowcontrol.RateLimiter
	recorder				events.Recorder
}

func NewOpenShiftControllerManagerOperator(targetImagePullSpec string, operatorConfigInformer operatorinformersv1.OpenShiftControllerManagerInformer, kubeInformersForOpenshiftControllerManager informers.SharedInformerFactory, operatorConfigClient operatorclientv1.OperatorV1Interface, kubeClient kubernetes.Interface, recorder events.Recorder) *OpenShiftControllerManagerOperator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &OpenShiftControllerManagerOperator{targetImagePullSpec: targetImagePullSpec, operatorConfigClient: operatorConfigClient, kubeClient: kubeClient, queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "KubeApiserverOperator"), rateLimiter: flowcontrol.NewTokenBucketRateLimiter(0.05, 4), recorder: recorder}
	operatorConfigInformer.Informer().AddEventHandler(c.eventHandler())
	kubeInformersForOpenshiftControllerManager.Core().V1().ConfigMaps().Informer().AddEventHandler(c.eventHandler())
	kubeInformersForOpenshiftControllerManager.Core().V1().ServiceAccounts().Informer().AddEventHandler(c.eventHandler())
	kubeInformersForOpenshiftControllerManager.Core().V1().Services().Informer().AddEventHandler(c.eventHandler())
	kubeInformersForOpenshiftControllerManager.Apps().V1().Deployments().Informer().AddEventHandler(c.eventHandler())
	kubeInformersForOpenshiftControllerManager.Core().V1().Namespaces().Informer().AddEventHandler(c.namespaceEventHandler())
	return c
}
func (c OpenShiftControllerManagerOperator) sync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	operatorConfig, err := c.operatorConfigClient.OpenShiftControllerManagers().Get("cluster", metav1.GetOptions{})
	if err != nil {
		return err
	}
	switch operatorConfig.Spec.ManagementState {
	case operatorapiv1.Unmanaged:
		originalOperatorConfig := operatorConfig.DeepCopy()
		v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorapiv1.OperatorCondition{Type: operatorapiv1.OperatorStatusTypeAvailable, Status: operatorapiv1.ConditionUnknown, Reason: "Unmanaged", Message: "the controller manager is in an unmanaged state, therefore its availability is unknown."})
		v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorapiv1.OperatorCondition{Type: operatorapiv1.OperatorStatusTypeProgressing, Status: operatorapiv1.ConditionFalse, Reason: "Unmanaged", Message: "the controller manager is in an unmanaged state, therefore no changes are being applied."})
		v1helpers.SetOperatorCondition(&operatorConfig.Status.Conditions, operatorapiv1.OperatorCondition{Type: operatorapiv1.OperatorStatusTypeFailing, Status: operatorapiv1.ConditionFalse, Reason: "Unmanaged", Message: "the controller manager is in an unmanaged state, therefore no operator actions are failing."})
		if !equality.Semantic.DeepEqual(operatorConfig.Status, originalOperatorConfig.Status) {
			if _, err := c.operatorConfigClient.OpenShiftControllerManagers().UpdateStatus(operatorConfig); err != nil {
				return err
			}
		}
		return nil
	case operatorapiv1.Removed:
		if err := c.kubeClient.CoreV1().Namespaces().Delete(util.TargetNamespace, nil); err != nil && !apierrors.IsNotFound(err) {
			return err
		}
		return nil
	}
	forceRequeue, err := syncOpenShiftControllerManager_v311_00_to_latest(c, operatorConfig)
	if forceRequeue && err != nil {
		c.queue.AddRateLimited(workQueueKey)
	}
	return err
}
func (c *OpenShiftControllerManagerOperator) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()
	klog.Infof("Starting OpenShiftControllerManagerOperator")
	defer klog.Infof("Shutting down OpenShiftControllerManagerOperator")
	go wait.Until(c.runWorker, time.Second, stopCh)
	<-stopCh
}
func (c *OpenShiftControllerManagerOperator) runWorker() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for c.processNextWorkItem() {
	}
}
func (c *OpenShiftControllerManagerOperator) processNextWorkItem() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dsKey, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(dsKey)
	c.rateLimiter.Accept()
	err := c.sync()
	if err == nil {
		c.queue.Forget(dsKey)
		return true
	}
	utilruntime.HandleError(fmt.Errorf("%v failed with : %v", dsKey, err))
	c.queue.AddRateLimited(dsKey)
	return true
}
func (c *OpenShiftControllerManagerOperator) eventHandler() cache.ResourceEventHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		c.queue.Add(workQueueKey)
	}, UpdateFunc: func(old, new interface{}) {
		c.queue.Add(workQueueKey)
	}, DeleteFunc: func(obj interface{}) {
		c.queue.Add(workQueueKey)
	}}
}
func (c *OpenShiftControllerManagerOperator) namespaceEventHandler() cache.ResourceEventHandler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.ResourceEventHandlerFuncs{AddFunc: func(obj interface{}) {
		ns, ok := obj.(*corev1.Namespace)
		if !ok {
			c.queue.Add(workQueueKey)
		}
		if ns.Name == util.TargetNamespace {
			c.queue.Add(workQueueKey)
		}
	}, UpdateFunc: func(old, new interface{}) {
		ns, ok := old.(*corev1.Namespace)
		if !ok {
			c.queue.Add(workQueueKey)
		}
		if ns.Name == util.TargetNamespace {
			c.queue.Add(workQueueKey)
		}
	}, DeleteFunc: func(obj interface{}) {
		ns, ok := obj.(*corev1.Namespace)
		if !ok {
			tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
			if !ok {
				utilruntime.HandleError(fmt.Errorf("couldn't get object from tombstone %#v", obj))
				return
			}
			ns, ok = tombstone.Obj.(*corev1.Namespace)
			if !ok {
				utilruntime.HandleError(fmt.Errorf("tombstone contained object that is not a Namespace %#v", obj))
				return
			}
		}
		if ns.Name == util.TargetNamespace {
			c.queue.Add(workQueueKey)
		}
	}}
}
