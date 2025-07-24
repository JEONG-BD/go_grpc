.PHONY: gen clean run server client test
gen:
	#protoc --proto_path=protoc --go_out=protoc/pb --go-grpc_out=protoc/pb protoc/*.proto
	#protoc --go_out=../pb --go-grpc_out=../pb *.proto     
	protoc --proto_path=protoc --go_out=pb --go-grpc_out=pb protoc/*.proto
clean:
	rm -f protoc/pb/*.go
run:
	#go run main.go

server:
	go run cmd/server/main.go -port=8080

client: 
	go run cmd/client/main.go -address 0.0.0.0:8080 

test:
	go test -cover -race ./...