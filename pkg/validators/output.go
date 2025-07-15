package validators

import (
    "encoding/json"
    "fmt"
    "strings"

    "github.com/fatih/color"
)

func PrintResults(results []ValidationResult, format, failOn string) int {
    exitCode := 0
    threshold := Severity(strings.ToLower(failOn))
    if format == "json" {
        data, _ := json.MarshalIndent(results, "", "  ")
        fmt.Println(string(data))
    } else {
        for _, r := range results {
            line := fmt.Sprintf("%-10s %-20s %-30s %-8s %s",
                r.Kind, r.Name, r.Rule, r.Status, r.Message)
            switch r.Status {
            case "FAIL":
                if r.Severity >= threshold {
                    color.Red(line)
                    exitCode = 1
                } else {
                    color.Yellow(line)
                }
            case "SKIPPED":
                color.Cyan(line)
            default:
                    fmt.Println(line)
            }
        }
    }
    return exitCode
}
