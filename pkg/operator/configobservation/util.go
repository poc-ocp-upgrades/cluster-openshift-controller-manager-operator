package configobservation

import (
	"bytes"
	"encoding/json"
	"strings"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ObserveField(observedConfig map[string]interface{}, val interface{}, fieldName string, skipIfEmpty bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nestedFields := strings.Split(fieldName, ".")
	if val == nil {
		return nil
	}
	var err error
	switch v := val.(type) {
	case int64, bool:
		err = unstructured.SetNestedField(observedConfig, v, nestedFields...)
	case string:
		if skipIfEmpty && len(v) == 0 {
			return nil
		}
		err = unstructured.SetNestedField(observedConfig, v, nestedFields...)
	case []interface{}:
		if skipIfEmpty && len(v) == 0 {
			return nil
		}
		err = unstructured.SetNestedSlice(observedConfig, v, nestedFields...)
	case map[string]string:
		if skipIfEmpty && len(v) == 0 {
			return nil
		}
		err = unstructured.SetNestedStringMap(observedConfig, v, nestedFields...)
	case map[string]interface{}:
		if skipIfEmpty && len(v) == 0 {
			return nil
		}
		err = unstructured.SetNestedMap(observedConfig, v, nestedFields...)
	default:
		data, err := ConvertJSON(v)
		if err != nil {
			return err
		}
		if skipIfEmpty && data == nil {
			return nil
		}
		err = unstructured.SetNestedField(observedConfig, data, nestedFields...)
	}
	return err
}
func ConvertJSON(o interface{}) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if o == nil {
		return nil, nil
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(o); err != nil {
		return nil, err
	}
	ret := []interface{}{}
	if err := json.NewDecoder(buf).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}
