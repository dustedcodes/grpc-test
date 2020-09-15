package main

import (
	"context"
	"log"

	"github.com/dustedcodes/grpc-test/server/hello"
	"google.golang.org/grpc"
)

func main() {
	addr := ":9000"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := hello.NewHelloClient(conn)
	res, err := client.SayHello(context.Background(), &hello.HelloRequest{Name: "Hello World!"})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Message)
}
