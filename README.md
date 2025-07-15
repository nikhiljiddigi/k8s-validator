# K8s Validator CLI Tool (MVP)

A CLI tool for validating Kubernetes manifests (YAML or Helm output) using a pluggable rule engine, with flexible exemption support.

## Features

- Rule-based validation of K8s objects
- Multi-container support with per-container validation
- Rule-centric `.exemptions.yaml` format
- Supports `global`, `kinds`, `files`, `namespaces`, and container exemptions
- Outputs in both table and JSON format
- Severity levels for rules (error, warning, info)
- Namespace-level exemptions
- Human-readable rule names instead of rule IDs

## How to Run

```bash
make run
```

Ensure you have Go installed and run `go mod tidy` first.
