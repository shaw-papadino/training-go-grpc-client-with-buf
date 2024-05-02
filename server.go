package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
    pb "training-go-grpc-client-with-buf/gen/go" // インポートパスを適切に調整してください
)

// server is used to implement geo.GreeterServer.
type server struct {
    pb.UnimplementedGreeterServer
}

// SayHello implements geo.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
