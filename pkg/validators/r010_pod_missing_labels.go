package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type PodMissingLabels struct{}

func (r PodMissingLabels) Name() string         { return "Missing Pod Labels" }
func (r PodMissingLabels) Severity() Severity    { return Warning }
func (r PodMissingLabels) AllowExemption() bool { return true }

func (r PodMissingLabels) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Pod" { return nil }
    if len(obj.GetLabels()) == 0 {
        return []ValidationResult{{
            Kind:      "Pod",
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
