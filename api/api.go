package api

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vchakoshy/graphdb/graph"
	"github.com/vchakoshy/graphdb/service"
	"google.golang.org/grpc"
)

var graphClient *graph.Graph

func Run() {
	graphClient = graph.NewGraph()

	// rest for test :D
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "This is Graph DB")
	})

	log.Println("Rest api Listening on: http://127.0.0.1:8081")

	go r.Run("0.0.0.0:8081")

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	service.SetGraphClient(graphClient)
	service.RegisterGraphdbServer(grpcServer, service.ImplementedGraphdbServer{})

	log.Println("Listening on: http://127.0.0.1:8080")
	grpcServer.Serve(lis)

}
