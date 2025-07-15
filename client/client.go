package main

import (
	"context"
	proto "grpc/protoc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//fmt.Println(err, conn)
	if err != nil {
		log.Panicln(err)
		panic(err)
	}

	client := proto.NewExampleClient(conn)

	req := &proto.HelloRequest{SomeString: "Hello from client"}

	res, err := client.ServerReply(context.TODO(), req)
	if err != nil {
		log.Panicln(err)
		panic(err)
	}

	log.Println(res)
}
