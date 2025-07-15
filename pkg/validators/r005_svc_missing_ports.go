package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type SvcMissingPorts struct{}

func (r SvcMissingPorts) Name() string         { return "Missing Service Ports" }
func (r SvcMissingPorts) Severity() Severity    { return Error }
func (r SvcMissingPorts) AllowExemption() bool { return true }

func (r SvcMissingPorts) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Service" { return nil }
    ports, found, _ := unstructured.NestedSlice(obj.Object, "spec", "ports")
    if !found || len(ports) == 0 {
        return []ValidationResult{{
            Kind:      "Service",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "spec.ports is missing or empty",
        }}
    }
    return nil
}
