package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type DeployStrategy struct{}

func (r DeployStrategy) Name() string         { return "Non-RollingUpdate Strategy" }
func (r DeployStrategy) Severity() Severity    { return Info }
func (r DeployStrategy) AllowExemption() bool { return true }

func (r DeployStrategy) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Deployment" { return nil }
    strat, _, _ := unstructured.NestedString(obj.Object, "spec", "strategy", "type")
    if strat != "RollingUpdate" {
        return []ValidationResult{{
            Kind:      "Deployment",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "strategy.type is not RollingUpdate",
        }}
    }
    return nil
}
