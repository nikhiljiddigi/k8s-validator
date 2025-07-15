package exemptions

import (
    "os"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

type RuleExemptions struct {
    Global     bool                `yaml:"global"`
    Kinds      []string            `yaml:"kinds"`
    Files      []string            `yaml:"files"`
    Namespaces []string            `yaml:"namespaces"`
    Containers map[string][]string `yaml:"containers"`
}

type ExemptConfig struct {
    Rules map[string]RuleExemptions `yaml:"rules"`
}

func LoadExemptions(path string) (*ExemptConfig, error) {
    cfg := &ExemptConfig{}
    if path == "" {
        return cfg, nil
    }
    b, err := os.ReadFile(filepath.Clean(path))
    if err != nil {
        return nil, err
    }
    if err := yaml.Unmarshal(b, cfg); err != nil {
        return nil, err
    }
    return cfg, nil
}
