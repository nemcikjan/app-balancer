package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/nemcikjan/app-balancer/grpc"
	"github.com/nemcikjan/app-balancer/k8s"
	"google.golang.org/grpc"
)

var (
	port         = flag.Int("port", 50053, "The server port")
	k8sClient, _ = k8s.NewK8sAppClient()
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAppBalancerServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) UpdateSpeed(ctx context.Context, in *pb.UpdatePodSpeed) (*pb.Result, error) {
	log.Printf("Received: %v, %f", in.GetId(), in.GetSpeed())
	k8sClient.UpdateApp(in.GetId(), in.GetSpeed())
	return &pb.Result{Message: "OK"}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) CreateApp(ctx context.Context, in *pb.CreatePodApp) (*pb.Result, error) {
	log.Printf("Received: %v", in.GetId())
	k8sClient.CreateApp(in.GetId())
	return &pb.Result{Message: "OK"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAppBalancerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
