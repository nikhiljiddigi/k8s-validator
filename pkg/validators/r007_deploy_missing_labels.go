package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type DeployMissingLabels struct{}

func (r DeployMissingLabels) Name() string         { return "Missing Deployment Labels" }
func (r DeployMissingLabels) Severity() Severity    { return Warning }
func (r DeployMissingLabels) AllowExemption() bool { return true }

func (r DeployMissingLabels) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Deployment" { return nil }
    if len(obj.GetLabels()) == 0 {
        return []ValidationResult{{
            Kind:      "Deployment",
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
