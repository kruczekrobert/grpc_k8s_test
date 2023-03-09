package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"grpc_k8s/compute"
	"log"
	"net"
)

var computeServiceCmd = &cobra.Command{
	Use:   "compute_service",
	Short: "run compute service",
	Run:   ComputeServiceCommandRun,
}

func ComputeServiceCommandRun(cmd *cobra.Command, args []string) {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	fmt.Println("Serve...")
	compute.RegisterAddServiceServer(srv, &compute.Service{})
	if err = srv.Serve(listener); err != nil {
		log.Fatalf("failed to servce: %v", err)
	}
}
