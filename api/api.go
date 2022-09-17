package api

import (
	"fmt"
	"log"
	"net"

	"github.com/vchakoshy/graphdb/graph"
	"github.com/vchakoshy/graphdb/service"
	"google.golang.org/grpc"
)

var graphClient *graph.Graph

func Run() {
	graphClient = graph.NewGraph()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
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
