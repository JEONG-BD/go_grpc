package main

import (
	"context"
	proto "grpc/protoc"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	client = proto.NewExampleClient(conn)
	r := gin.Default()
	r.Use(gindump.Dump())
	r.GET("/send", clientConnectionServer)
	r.Run(":9001")
}

func clientConnectionServer(c *gin.Context) {
	stream, err := client.ServerReply(context.TODO())

	if err != nil {
		log.Println(err)
		return
	}

	send, receive := 0, 0

	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloRequest{SomeString: "Test"})
		if err != nil {
			log.Println(err)
			return
		}
		send++
	}

	if err := stream.CloseSend(); err != nil {
		log.Println(err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}
		log.Println("Server message: -", message)
		receive++
	}

	c.JSON(http.StatusOK, gin.H{
		"message sent":    send,
		"message reveive": receive,
	})

}
