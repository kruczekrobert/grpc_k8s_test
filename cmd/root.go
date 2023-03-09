package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "grpc_k8s",
	Short: "grpc_k8s",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(computeServiceCmd)
	rootCmd.AddCommand(apiCommand)
}
