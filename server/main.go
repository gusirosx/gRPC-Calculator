package main

import (
	"context"
	pb "gRPC/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// essa estrutura deve ser a primeira coisa a ser criada para que ela possa implementar a interface que foi gerada pelo protobuf
type server struct {
	pb.UnimplementedAddServiceServer
}

//o server type deve ter que implementar ambas as funções registradas no proto

func main() {

	//Begin listening on the port specified
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer() //Variable 'srv' is invoking the new server function from the grpc module
	//After initialize this new server we're going to register the server as a new grpc service
	pb.RegisterAddServiceServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
	// return s.Serve(lis) //Call the server

}
func (s *server) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &pb.Response{Result: result}, nil
}

//http://localhost:8080/add/125/125
//http://localhost:8080/Multiply/125/125
