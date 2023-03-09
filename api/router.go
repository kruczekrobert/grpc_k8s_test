package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"grpc_k8s/compute"
	"log"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", handlerFunc)

	return r
}

func handlerFunc(ctx *gin.Context) {
	conn, err := grpc.Dial("compute-service:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial Failed: %v", err)
	}
	addClient := compute.NewAddServiceClient(conn)

	cCtx, cancel := context.WithTimeout(context.TODO(), time.Minute)
	defer cancel()

	req := &compute.AddRequest{A: 1, B: 2}
	resp, err := addClient.Compute(cCtx, req)
	if err != nil {
		log.Println("[10245245]", err)
		ctx.Status(500)
		return
	}

	ctx.JSON(200, gin.H{"result": resp.Result})
}
