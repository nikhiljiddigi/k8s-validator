package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type SvcLBNoSourceRanges struct{}

func (r SvcLBNoSourceRanges) Name() string         { return "LoadBalancer Without SourceRanges" }
func (r SvcLBNoSourceRanges) Severity() Severity    { return Warning }
func (r SvcLBNoSourceRanges) AllowExemption() bool { return true }

func (r SvcLBNoSourceRanges) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Service" { return nil }
    typ, _, _ := unstructured.NestedString(obj.Object, "spec", "type")
    if typ == "LoadBalancer" {
        _, found, _ := unstructured.NestedSlice(obj.Object, "spec", "loadBalancerSourceRanges")
        if !found {
            return []ValidationResult{{
                Kind:      "Service",
                Name:      obj.GetName(),
                Namespace: obj.GetNamespace(),
                Rule:      r.Name(),
                Severity:  r.Severity(),
                Status:    "FAIL",
                Message:   "type=LoadBalancer without loadBalancerSourceRanges",
            }}
        }
    }
    return nil
}
