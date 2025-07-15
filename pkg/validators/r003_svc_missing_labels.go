package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type SvcMissingLabels struct{}

func (r SvcMissingLabels) Name() string         { return "Missing Service Labels" }
func (r SvcMissingLabels) Severity() Severity    { return Warning }
func (r SvcMissingLabels) AllowExemption() bool { return true }

func (r SvcMissingLabels) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Service" { return nil }
    if len(obj.GetLabels()) == 0 {
        return []ValidationResult{{
            Kind:      "Service",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "metadata.labels is empty",
        }}
    }
    return nil
}
