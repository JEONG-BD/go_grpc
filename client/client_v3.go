package main

import (
	"context"
	proto "grpc/protoc"
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
	req := []*proto.HelloRequest{
		{SomeString: "Rquest1"},
		{SomeString: "Rquest2"},
		{SomeString: "Rquest3"},
	}
	stream, err := client.ServerReply(context.TODO())

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for _, re := range req {
		err := stream.Send(re)

		if err != nil {
			log.Println("request not fullfill")
			return
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("There is some error occured", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message_count": res,
	})

}
