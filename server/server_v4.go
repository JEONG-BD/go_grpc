package main

import (
	proto "grpc/protoc"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedStreamUploadServer
}

func main() {
	log.Println("Server Start")
	listener, tcpErr := net.Listen("tcp", ":9000")

	if tcpErr != nil {
		log.Println(tcpErr)
		return
	}

	srv := grpc.NewServer()
	proto.RegisterStreamUploadServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s server) Upload(stream proto.StreamUpload_UploadServer) error {
	var fileBytes []byte
	var fileSize int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		chunks := req.GetChunks()
		fileBytes = append(fileBytes, chunks...)
		fileSize += int64(len(chunks))
	}

	f, err := os.Create("./abc.bin")
	if err != nil {
		log.Println(err)
		return err
	}

	defer f.Close()
	_, err2 := f.Write(fileBytes)
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	return stream.SendAndClose(&proto.UploadResponse{FileSize: fileSize, Message: "File Write successed"})

}
