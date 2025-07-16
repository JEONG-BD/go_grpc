package main

import (
	proto "grpc/protoc"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	log.Println("Server start")
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		log.Println(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		log.Panic(e)
	}
}

func (s *server) ServerReply(strem proto.Example_ServerReplyServer) error {
	for i := 0; i < 10; i++ {
		err := strem.Send(&proto.HelloResponse{Reply: "message " + strconv.Itoa(i) + " from Server"})

		if err != nil {
			log.Println(err)
			panic(err)
		}
	}

	for {
		req, err := strem.Recv()
		if err == io.EOF {
			break
		}
		log.Println(req.SomeString)

	}
	return nil

}
