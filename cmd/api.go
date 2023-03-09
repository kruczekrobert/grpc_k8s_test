package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"grpc_k8s/api"
)

var apiCommand = &cobra.Command{
	Use:   "api",
	Short: "run api",
	Run:   ApiCommandRun,
}

func ApiCommandRun(cmd *cobra.Command, args []string) {
	r := api.NewRouter()
	fmt.Println("Api serve...")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
