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
		panic(tcpErr)
	}

	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		log.Println(e)
		panic(e)
	}
}

func (s *server) ServerReply(strem proto.Example_ServerReplyServer) error {
	total := 0
	log.Println("ServerStream")
	for {
		request, err := strem.Recv()
		if err == io.EOF {
			return strem.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa(total),
			})
		}

		if err != nil {
			log.Println(err)
			return err
		}

		total++
		log.Println(request)
	}

}
