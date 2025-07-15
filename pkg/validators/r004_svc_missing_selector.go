package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type SvcMissingSelector struct{}

func (r SvcMissingSelector) Name() string         { return "Missing Service Selector" }
func (r SvcMissingSelector) Severity() Severity    { return Error }
func (r SvcMissingSelector) AllowExemption() bool { return true }

func (r SvcMissingSelector) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Service" { return nil }
    sel, found, _ := unstructured.NestedMap(obj.Object, "spec", "selector")
    if !found || len(sel) == 0 {
        return []ValidationResult{{
            Kind:      "Service",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "spec.selector is missing or empty",
        }}
    }
    return nil
}
