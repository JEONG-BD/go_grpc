package main

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// global variable

var client proto.ExampleClient

func main() {
	// Connection to internal grpc server
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println(err, conn)
	if err != nil {
		log.Panicln(err)
		panic(err)
	}

	client = proto.NewExampleClient(conn)

	r := gin.Default()
	r.Use(gindump.Dump())
	r.GET("/sent-message-to-server/:message", clientConnectionServer)
	r.Run(":9001")
	req := &proto.HelloRequest{SomeString: "Hello from client"}

	res, err := client.ServerReply(context.TODO(), req)
	if err != nil {
		log.Panicln(err)
		panic(err)
	}

	log.Println(res)
}

func clientConnectionServer(c *gin.Context) {
	message := c.Param("message")
	req := &proto.HelloRequest{SomeString: message}
	client.ServerReply(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent sucessfully to server : " + message,
	})
}
