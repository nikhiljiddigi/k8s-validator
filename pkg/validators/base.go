package validators

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type Severity string

const (
    Error   Severity = "error"
    Warning Severity = "warning"
    Info    Severity = "info"
)

type ValidationResult struct {
    Kind      string   `json:"kind"`
    Name      string   `json:"name"`
    Namespace string   `json:"namespace"`
    Container string   `json:"container,omitempty"`
    Rule      string   `json:"rule"`
    Severity  Severity `json:"severity"`
    Status    string   `json:"status"`
    Message   string   `json:"message"`
}

type Rule interface {
    Name() string
    Severity() Severity
    AllowExemption() bool
    Validate(unstructured.Unstructured) []ValidationResult
}
