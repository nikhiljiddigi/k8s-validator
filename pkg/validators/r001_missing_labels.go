package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type MissingLabels struct{}

func (r MissingLabels) Name() string         { return "Missing Labels" }
func (r MissingLabels) Severity() Severity    { return Warning }
func (r MissingLabels) AllowExemption() bool { return true }

func (r MissingLabels) Validate(obj unstructured.Unstructured) []ValidationResult {
    if len(obj.GetLabels()) == 0 {
        return []ValidationResult{{
            Kind:      obj.GetKind(),
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
