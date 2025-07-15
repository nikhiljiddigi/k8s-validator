package validators

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type MissingLimits struct{}

func (r MissingLimits) Name() string         { return "Missing Resource Limits" }
func (r MissingLimits) Severity() Severity   { return Error }
func (r MissingLimits) AllowExemption() bool { return true }

func (r MissingLimits) Validate(obj unstructured.Unstructured) []ValidationResult {
	kind := obj.GetKind()
	var specMap map[string]interface{}
	var found bool
	var err error

	if kind == "Pod" {
		specMap, found, err = unstructured.NestedMap(obj.Object, "spec")
	} else {
		specMap, found, err = unstructured.NestedMap(obj.Object, "spec", "template", "spec")
	}
	if err != nil || !found {
		return nil
	}

	containers, found, err := unstructured.NestedSlice(specMap, "containers")
	if err != nil || !found {
		return nil
	}

	for _, c := range containers {
		cm := c.(map[string]interface{})
		name := cm["name"].(string)
		limits, found, _ := unstructured.NestedMap(cm, "resources", "limits")
		if !found || len(limits) == 0 {
			return []ValidationResult{{
				Kind:      kind,
				Name:      obj.GetName(),
				Namespace: obj.GetNamespace(),
				Container: name,
				Rule:      r.Name(),
				Severity:  r.Severity(),
				Status:    "FAIL",
				Message:   "resources.limits is missing",
			}}
		}
	}
	return nil
}
