package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type PodRunAsNonRoot struct{}

func (r PodRunAsNonRoot) Name() string         { return "RunAsNonRoot Not Set" }
func (r PodRunAsNonRoot) Severity() Severity    { return Error }
func (r PodRunAsNonRoot) AllowExemption() bool { return true }

func (r PodRunAsNonRoot) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Pod" { return nil }
    sc, found, _ := unstructured.NestedMap(obj.Object, "spec", "securityContext")
    if !found {
        return []ValidationResult{{
            Kind:      "Pod",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "spec.securityContext missing",
        }}
    }
    if run, ok := sc["runAsNonRoot"].(bool); !ok || !run {
        return []ValidationResult{{
            Kind:      "Pod",
            Name:      obj.GetName(),
            Namespace: obj.GetNamespace(),
            Rule:      r.Name(),
            Severity:  r.Severity(),
            Status:    "FAIL",
            Message:   "securityContext.runAsNonRoot is not true",
        }}
    }
    return nil
}
