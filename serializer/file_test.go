package serializer_test

import (
	"fmt"
	"grpc/pb"
	"grpc/sample"
	"grpc/serializer"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

// go get github.com/stretchr/testify
// go test -v ./grpc/serializer
func TestFileSerializer(t *testing.T) {
	t.Parallel()

	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)
	binaryFile := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

}

func TestFileSerializer2(t *testing.T) {
	t.Parallel()

	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	binaryFile := "../tmp/laptop.bin"

	err := os.MkdirAll(filepath.Dir(binaryFile), os.ModePerm)
	require.NoError(t, err)

	laptop1 := sample.NewLaptop()
	err = serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}

func TestJsonFileSerializer(t *testing.T) {
	t.Parallel()
	laptop1 := sample.NewLaptop()
	jsonFile := "../tmp/laptop.json"
	err := serializer.WriteProtobufToJsonFile(laptop1, jsonFile)
	require.NoError(t, err)
}
