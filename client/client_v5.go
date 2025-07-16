package main

import (
	"context"
	"fmt"
	proto "grpc/protoc"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.StreamUploadClient

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	client = proto.NewStreamUploadClient(conn)

	makeMb := 1024 * 1024 * 1024
	createDummyFile("./1GB.bin", makeMb)
	sendMb := 1024 * 1024 * 2
	uploadStreamFile("./1GB.bin", sendMb)

}

func createDummyFile(path string, sizeInBytes int) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("파일 생성 실패: %v", err)
	}
	defer file.Close()

	err = file.Truncate(int64(sizeInBytes))
	if err != nil {
		log.Fatalf("파일 크기 설정 실패: %v", err)
	}
}

func uploadStreamFile(path string, batchSize int) {
	t := time.Now()
	file, err := os.Open(path)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	buf := make([]byte, batchSize)
	batchNumber := 1
	stream, err := client.Upload(context.TODO())

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		chunk := buf[:num]

		if err := stream.Send(&proto.UploadRequest{FilePath: path, Chunks: chunk}); err != nil {
			fmt.Println(err)
			return
		}

		log.Printf("Senf - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(time.Since(t))
	log.Printf("Sent - %v bytes - %s \n", res.GetFileSize(), res.GetMessage())

}
