package main

import (
	"context"
	"flag"
	"fmt"
	"grpc/pb"
	"grpc/sample"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Hello world from client")
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)
	con, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can not dial server : ", err)
	}
	laptopClient := pb.NewLaptopServiceClient(con)

	laptop := sample.NewLaptop()

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)

	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Panicln("laptop already exists")
		} else {
			log.Fatal("can not create laptop")
		}
		return
	}

	log.Printf("created laptop with id: %s", res.Id)

}
