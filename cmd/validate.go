package cmd

import (
	"fmt"
	"os"

	"k8s-validator/pkg/exemptions"
	"k8s-validator/pkg/loader"
	"k8s-validator/pkg/validators"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	inputDir     string
	chartDir     string
	valuesFile   string
	exemptFile   string
	outputFormat string
	failOn       string
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Run validations against K8s manifests or Helm charts",
	Run: func(cmd *cobra.Command, args []string) {
		// Must specify input
		if inputDir == "" && chartDir == "" {
			fmt.Println("Error: you must provide either -f (folder) or -c (chart)")
			os.Exit(1)
		}

		// Load K8s objects
		var objs []unstructured.Unstructured
		var err error
		if inputDir != "" {
			objs, err = loader.LoadYAMLFolder(inputDir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "YAML load error: %v\n", err)
				os.Exit(1)
			}
		}
		if chartDir != "" {
			helmObjs, e := loader.RenderHelmChart(chartDir, valuesFile)
			if e != nil {
				fmt.Fprintf(os.Stderr, "Helm render error: %v\n", e)
				os.Exit(1)
			}
			objs = append(objs, helmObjs...)
		}

		// Load exemptions
		ex, err := exemptions.LoadExemptions(exemptFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Exemptions load error: %v\n", err)
			os.Exit(1)
		}

		// Run all rules
		results := validators.RunAll(objs, ex)

		// Print and exit according to --fail-on
		code := validators.PrintResults(results, outputFormat, failOn)
		os.Exit(code)
	},
}

func init() {
	// Register flags (with shorthands for folder and chart)
	validateCmd.Flags().StringVarP(&inputDir, "folder", "f", "", "Path to folder of YAML manifests")
	validateCmd.Flags().StringVarP(&chartDir, "chart", "c", "", "Path to Helm chart directory")
	validateCmd.Flags().StringVar(&valuesFile, "values", "", "Path to Helm values.yaml")
	validateCmd.Flags().StringVar(&exemptFile, "exemptions", "", "Path to rule-centric exemptions YAML")
	validateCmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format: table or json")
	validateCmd.Flags().StringVar(&failOn, "fail-on", "error", "Fail on severity: info, warning, or error (default: error)")

	// Attach validate command to the root
	rootCmd.AddCommand(validateCmd)
}
