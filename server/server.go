package main

import (
	"log"
	"net"

	"github.com/dustedcodes/grpc-test/server/hello"
	"google.golang.org/grpc"
)

func main() {
	addr := ":9000"
	log.Println("Listening on " + addr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	hello.RegisterHelloServer(grpcServer, &hello.Service{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
