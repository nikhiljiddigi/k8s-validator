package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type DeployMissingProbes struct{}

func (r DeployMissingProbes) Name() string         { return "Missing Readiness/Liveness Probe" }
func (r DeployMissingProbes) Severity() Severity    { return Warning }
func (r DeployMissingProbes) AllowExemption() bool { return true }

func (r DeployMissingProbes) Validate(obj unstructured.Unstructured) []ValidationResult {
    if obj.GetKind() != "Deployment" { return nil }
    podSpec, found, _ := unstructured.NestedMap(obj.Object, "spec", "template", "spec")
    if !found { return nil }
    containers, _, _ := unstructured.NestedSlice(podSpec, "containers")
    var results []ValidationResult
    for _, c := range containers {
        cm := c.(map[string]interface{})
        name := cm["name"].(string)
        if _, f1, _ := unstructured.NestedMap(cm, "readinessProbe"); !f1 {
            results = append(results, ValidationResult{
                Kind:      "Deployment",
                Name:      obj.GetName(),
                Namespace: obj.GetNamespace(),
                Container: name,
                Rule:      r.Name(),
                Severity:  r.Severity(),
                Status:    "FAIL",
                Message:   "readinessProbe missing",
            })
        }
        if _, f2, _ := unstructured.NestedMap(cm, "livenessProbe"); !f2 {
            results = append(results, ValidationResult{
                Kind:      "Deployment",
                Name:      obj.GetName(),
                Namespace: obj.GetNamespace(),
                Container: name,
                Rule:      r.Name(),
                Severity:  r.Severity(),
                Status:    "FAIL",
                Message:   "livenessProbe missing",
            })
        }
    }
    return results
}
