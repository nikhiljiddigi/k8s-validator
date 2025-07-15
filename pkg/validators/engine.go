package validators

import (
	"path/filepath"
	"strings"

	"k8s-validator/pkg/exemptions"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func RunAll(objs []unstructured.Unstructured, ex *exemptions.ExemptConfig) []ValidationResult {
	// 1) Define your kind→rules map
	kindRules := map[string][]Rule{
		"Service": {
			SvcMissingLabels{},
			SvcMissingSelector{},
			SvcMissingPorts{},
			SvcLBNoSourceRanges{},
		},
		"Deployment": {
			DeployMissingLabels{},
			DeployMissingProbes{},
			DeployStrategy{},
		},
		"Pod": {
			PodMissingLabels{},
			MissingLabels{}, // your existing label rule
			MissingLimits{}, // your existing limits rule
			PodHostNetwork{},
			PodRunAsNonRoot{},
		},
	}

	var out []ValidationResult

	// 2) Iterate objects, only their own rules
	for _, obj := range objs {
		rules := kindRules[obj.GetKind()]
		for _, rule := range rules {
			name := rule.Name()

			// exemption check (unchanged)
			if shouldSkip(obj, name, ex) {
				out = append(out, ValidationResult{
					Kind:      obj.GetKind(),
					Name:      obj.GetName(),
					Namespace: obj.GetNamespace(),
					Rule:      name,
					Severity:  rule.Severity(),
					Status:    "SKIPPED",
					Message:   "exempted by config",
				})
				continue
			}

			// run only relevant rules
			results := rule.Validate(obj)
			if len(results) > 0 {
				out = append(out, results...)
			} else {
				// PASS only for rules that ran
				out = append(out, ValidationResult{
					Kind:      obj.GetKind(),
					Name:      obj.GetName(),
					Namespace: obj.GetNamespace(),
					Rule:      name,
					Severity:  rule.Severity(),
					Status:    "PASS",
					Message:   "passed",
				})
			}
		}
	}

	return out
}

func shouldSkip(obj unstructured.Unstructured, ruleName string, ex *exemptions.ExemptConfig) bool {
	cfg, ok := ex.Rules[ruleName]
	if !ok {
		return false
	}
	if cfg.Global {
		return true
	}
	for _, k := range cfg.Kinds {
		if strings.EqualFold(k, obj.GetKind()) {
			return true
		}
	}
	for _, f := range cfg.Files {
		if f == obj.GetName() {
			return true
		}
	}
	for _, ns := range cfg.Namespaces {
		if match, _ := filepath.Match(ns, obj.GetNamespace()); match {
			return true
		}
	}
	return false
}
