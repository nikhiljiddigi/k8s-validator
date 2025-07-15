// cmd/root.go
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-validate",
	Short: "Validate Kubernetes manifests and Helm charts",
}

// Execute runs the root command, called by main()
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// attach the validate subcommand
	rootCmd.AddCommand(validateCmd)
}
