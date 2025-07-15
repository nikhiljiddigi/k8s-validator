package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type PodHostNetwork struct{}

func (r PodHostNetwork) Name() string         { return "HostNetwork Enabled" }
func (r PodHostNetwork) Severity() Severity    { return Error }
func (r PodHostNetwork) AllowExemption() bool { return true }

func (r PodHostNetwork) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Pod" { return nil }
    if hn, found, _ := unstructured.NestedBool(obj.Object, "spec", "hostNetwork"); found && hn {
        return []ValidationResult{{
            Kind:      "Pod",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "spec.hostNetwork is true",
        }}
    }
    return nil
}
