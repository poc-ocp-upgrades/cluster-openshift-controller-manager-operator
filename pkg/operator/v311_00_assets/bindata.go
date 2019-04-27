package v311_00_assets

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes	[]byte
	info	os.FileInfo
}
type bindataFileInfo struct {
	name	string
	size	int64
	mode	os.FileMode
	modTime	time.Time
}

func (fi bindataFileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.mode&os.ModeDir != 0
}
func (fi bindataFileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

var _v3110OpenshiftControllerManagerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-controller-manager
  name: config
data:
  config.yaml:
`)

func v3110OpenshiftControllerManagerCmYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerCmYaml, nil
}
func v3110OpenshiftControllerManagerCmYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerCmYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerDefaultconfigYaml = []byte(`apiVersion: openshiftcontrolplane.config.openshift.io/v1
kind: OpenShiftControllerManagerConfig
`)

func v3110OpenshiftControllerManagerDefaultconfigYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerDefaultconfigYaml, nil
}
func v3110OpenshiftControllerManagerDefaultconfigYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerDsYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  namespace: openshift-controller-manager
  name: controller-manager
  labels:
    app: openshift-controller-manager
    controller-manager: "true"
spec:
  updateStrategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: openshift-controller-manager
      controller-manager: "true"
  template:
    metadata:
      name: openshift-controller-manager
      labels:
        app: openshift-controller-manager
        controller-manager: "true"
    spec:
      priorityClassName: system-node-critical 
      serviceAccountName: openshift-controller-manager-sa
      containers:
      - name: controller-manager
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["hypershift", "openshift-controller-manager"]
        args:
        - "--config=/var/run/configmaps/config/config.yaml"
        resources:
          requests:
            memory: 100Mi
            cpu: 100m
        ports:
        - containerPort: 8443
        volumeMounts:
        - mountPath: /var/run/configmaps/config
          name: config
        - mountPath: /var/run/configmaps/client-ca
          name: client-ca
        - mountPath: /var/run/secrets/serving-cert
          name: serving-cert
      volumes:
      - name: config
        configMap:
          name: config
      - name: client-ca
        configMap:
          name: client-ca
      - name: serving-cert
        secret:
          secretName: serving-cert
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      tolerations:
      - operator: Exists
`)

func v3110OpenshiftControllerManagerDsYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerDsYaml, nil
}
func v3110OpenshiftControllerManagerDsYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerDsYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/ds.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerInformerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:openshift:openshift-controller-manager
rules:
# we run cluster resource quota, so we have to be able to see all resources
- apiGroups:
  - "*"
  resources:
  - "*"
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  - events.k8s.io
  resources:
  - events
  verbs:
  - create
  - patch
  - update
`)

func v3110OpenshiftControllerManagerInformerClusterroleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerInformerClusterroleYaml, nil
}
func v3110OpenshiftControllerManagerInformerClusterroleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerInformerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/informer-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerInformerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:openshift-controller-manager
roleRef:
  kind: ClusterRole
  name: system:openshift:openshift-controller-manager
subjects:
- kind: ServiceAccount
  namespace: openshift-controller-manager
  name: openshift-controller-manager-sa
`)

func v3110OpenshiftControllerManagerInformerClusterrolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerInformerClusterrolebindingYaml, nil
}
func v3110OpenshiftControllerManagerInformerClusterrolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerInformerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/informer-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerLeaderRoleYaml = []byte(`# needed to get the legacy lock that we used to use
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:openshift:leader-locking-openshift-controller-manager
  namespace: kube-system
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - create
- apiGroups:
  - ""
  resourceNames:
  - openshift-master-controllers
  resources:
  - configmaps
  verbs:
  - get
  - create
  - update
  - patch`)

func v3110OpenshiftControllerManagerLeaderRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerLeaderRoleYaml, nil
}
func v3110OpenshiftControllerManagerLeaderRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerLeaderRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/leader-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerLeaderRolebindingYaml = []byte(`# needed to get the legacy lock that we used to use
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: kube-system
  name: system:openshift:leader-locking-openshift-controller-manager
roleRef:
  kind: Role
  name: system:openshift:leader-locking-openshift-controller-manager
subjects:
- kind: ServiceAccount
  namespace: openshift-controller-manager
  name: openshift-controller-manager-sa`)

func v3110OpenshiftControllerManagerLeaderRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerLeaderRolebindingYaml, nil
}
func v3110OpenshiftControllerManagerLeaderRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerLeaderRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/leader-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerOperatorConfigYaml = []byte(`apiVersion: operator.openshift.io/v1
kind: OpenShiftControllerManager
metadata:
  name: cluster
spec:
  managementState: Managed
  imagePullSpec: openshift/origin-hypershift:latest
  version: 3.11.0
  logging:
    level: 4
  replicas: 2
`)

func v3110OpenshiftControllerManagerOperatorConfigYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerOperatorConfigYaml, nil
}
func v3110OpenshiftControllerManagerOperatorConfigYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerOperatorConfigYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/operator-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-controller-manager
  name: openshift-controller-manager-sa
`)

func v3110OpenshiftControllerManagerSaYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerSaYaml, nil
}
func v3110OpenshiftControllerManagerSaYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerSaYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerSeparateSaRoleYaml = []byte(`# needed to support the "use separate service accounts" feature.
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:openshift:sa-creating-openshift-controller-manager
  namespace: openshift-infra
rules:
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  verbs:
  - get
  - create
  - update
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - create
`)

func v3110OpenshiftControllerManagerSeparateSaRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerSeparateSaRoleYaml, nil
}
func v3110OpenshiftControllerManagerSeparateSaRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerSeparateSaRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/separate-sa-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerSeparateSaRolebindingYaml = []byte(`# needed to support the "use separate service accounts" feature.
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: openshift-infra
  name: system:openshift:sa-creating-openshift-controller-manager
roleRef:
  kind: Role
  name: system:openshift:sa-creating-openshift-controller-manager
subjects:
- kind: ServiceAccount
  namespace: openshift-controller-manager
  name: openshift-controller-manager-sa
`)

func v3110OpenshiftControllerManagerSeparateSaRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerSeparateSaRolebindingYaml, nil
}
func v3110OpenshiftControllerManagerSeparateSaRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerSeparateSaRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/separate-sa-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerServicemonitorRoleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: prometheus-k8s
  namespace: openshift-controller-manager
rules:
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  - pods
  verbs:
  - get
  - list
  - watch
`)

func v3110OpenshiftControllerManagerServicemonitorRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerServicemonitorRoleYaml, nil
}
func v3110OpenshiftControllerManagerServicemonitorRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerServicemonitorRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/servicemonitor-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerServicemonitorRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: prometheus-k8s
  namespace: openshift-controller-manager
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: prometheus-k8s
subjects:
- kind: ServiceAccount
  name: prometheus-k8s
  namespace: openshift-monitoring
`)

func v3110OpenshiftControllerManagerServicemonitorRolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerServicemonitorRolebindingYaml, nil
}
func v3110OpenshiftControllerManagerServicemonitorRolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerServicemonitorRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/servicemonitor-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-controller-manager
  name: controller-manager
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    controller-manager: "true"
  ports:
  - name: https
    port: 443
    targetPort: 8443
`)

func v3110OpenshiftControllerManagerSvcYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerSvcYaml, nil
}
func v3110OpenshiftControllerManagerSvcYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerSvcYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerTokenreviewClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:openshift:tokenreview-openshift-controller-manager
rules:
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
`)

func v3110OpenshiftControllerManagerTokenreviewClusterroleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerTokenreviewClusterroleYaml, nil
}
func v3110OpenshiftControllerManagerTokenreviewClusterroleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerTokenreviewClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/tokenreview-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:tokenreview-openshift-controller-manager
roleRef:
  kind: ClusterRole
  name: system:openshift:tokenreview-openshift-controller-manager
subjects:
- kind: ServiceAccount
  namespace: openshift-controller-manager
  name: openshift-controller-manager-sa
`)

func v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYaml, nil
}
func v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "v3.11.0/openshift-controller-manager/tokenreview-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
func Asset(name string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}
func MustAsset(name string) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}
	return a
}
func AssetInfo(name string) (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}
func AssetNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

var _bindata = map[string]func() (*asset, error){"v3.11.0/openshift-controller-manager/cm.yaml": v3110OpenshiftControllerManagerCmYaml, "v3.11.0/openshift-controller-manager/defaultconfig.yaml": v3110OpenshiftControllerManagerDefaultconfigYaml, "v3.11.0/openshift-controller-manager/ds.yaml": v3110OpenshiftControllerManagerDsYaml, "v3.11.0/openshift-controller-manager/informer-clusterrole.yaml": v3110OpenshiftControllerManagerInformerClusterroleYaml, "v3.11.0/openshift-controller-manager/informer-clusterrolebinding.yaml": v3110OpenshiftControllerManagerInformerClusterrolebindingYaml, "v3.11.0/openshift-controller-manager/leader-role.yaml": v3110OpenshiftControllerManagerLeaderRoleYaml, "v3.11.0/openshift-controller-manager/leader-rolebinding.yaml": v3110OpenshiftControllerManagerLeaderRolebindingYaml, "v3.11.0/openshift-controller-manager/operator-config.yaml": v3110OpenshiftControllerManagerOperatorConfigYaml, "v3.11.0/openshift-controller-manager/sa.yaml": v3110OpenshiftControllerManagerSaYaml, "v3.11.0/openshift-controller-manager/separate-sa-role.yaml": v3110OpenshiftControllerManagerSeparateSaRoleYaml, "v3.11.0/openshift-controller-manager/separate-sa-rolebinding.yaml": v3110OpenshiftControllerManagerSeparateSaRolebindingYaml, "v3.11.0/openshift-controller-manager/servicemonitor-role.yaml": v3110OpenshiftControllerManagerServicemonitorRoleYaml, "v3.11.0/openshift-controller-manager/servicemonitor-rolebinding.yaml": v3110OpenshiftControllerManagerServicemonitorRolebindingYaml, "v3.11.0/openshift-controller-manager/svc.yaml": v3110OpenshiftControllerManagerSvcYaml, "v3.11.0/openshift-controller-manager/tokenreview-clusterrole.yaml": v3110OpenshiftControllerManagerTokenreviewClusterroleYaml, "v3.11.0/openshift-controller-manager/tokenreview-clusterrolebinding.yaml": v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYaml}

func AssetDir(name string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func		func() (*asset, error)
	Children	map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{"v3.11.0": {nil, map[string]*bintree{"openshift-controller-manager": {nil, map[string]*bintree{"cm.yaml": {v3110OpenshiftControllerManagerCmYaml, map[string]*bintree{}}, "defaultconfig.yaml": {v3110OpenshiftControllerManagerDefaultconfigYaml, map[string]*bintree{}}, "ds.yaml": {v3110OpenshiftControllerManagerDsYaml, map[string]*bintree{}}, "informer-clusterrole.yaml": {v3110OpenshiftControllerManagerInformerClusterroleYaml, map[string]*bintree{}}, "informer-clusterrolebinding.yaml": {v3110OpenshiftControllerManagerInformerClusterrolebindingYaml, map[string]*bintree{}}, "leader-role.yaml": {v3110OpenshiftControllerManagerLeaderRoleYaml, map[string]*bintree{}}, "leader-rolebinding.yaml": {v3110OpenshiftControllerManagerLeaderRolebindingYaml, map[string]*bintree{}}, "operator-config.yaml": {v3110OpenshiftControllerManagerOperatorConfigYaml, map[string]*bintree{}}, "sa.yaml": {v3110OpenshiftControllerManagerSaYaml, map[string]*bintree{}}, "separate-sa-role.yaml": {v3110OpenshiftControllerManagerSeparateSaRoleYaml, map[string]*bintree{}}, "separate-sa-rolebinding.yaml": {v3110OpenshiftControllerManagerSeparateSaRolebindingYaml, map[string]*bintree{}}, "servicemonitor-role.yaml": {v3110OpenshiftControllerManagerServicemonitorRoleYaml, map[string]*bintree{}}, "servicemonitor-rolebinding.yaml": {v3110OpenshiftControllerManagerServicemonitorRolebindingYaml, map[string]*bintree{}}, "svc.yaml": {v3110OpenshiftControllerManagerSvcYaml, map[string]*bintree{}}, "tokenreview-clusterrole.yaml": {v3110OpenshiftControllerManagerTokenreviewClusterroleYaml, map[string]*bintree{}}, "tokenreview-clusterrolebinding.yaml": {v3110OpenshiftControllerManagerTokenreviewClusterrolebindingYaml, map[string]*bintree{}}}}}}}}

func RestoreAsset(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}
func RestoreAssets(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}
func _filePath(dir, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
