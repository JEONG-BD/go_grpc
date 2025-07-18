gen:
	#protoc --proto_path=protoc --go_out=protoc/pb --go-grpc_out=protoc/pb protoc/*.proto
	#protoc --go_out=../pb --go-grpc_out=../pb *.proto     
	protoc --proto_path=protoc --go_out=pb --go-grpc_out=pb protoc/*.proto
clean:
	rm -f protoc/pb/*.go

run:
	go run main.go
