package serializer

//import "google.golang.org/grpc/encoding/proto"
import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {

	data, err := proto.Marshal(message)

	if err != nil {
		return fmt.Errorf("can not marchal proto message to binary : %w", err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("can not write binary data to file : %w", err)

	}
	return nil
}
func WriteProtobufToJsonFile(message proto.Message, filename string) error {

	data, err := ProtobufToJSON(message)

	if err != nil {
		return fmt.Errorf("can not marchal proto message to JSON : %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)

	if err != nil {
		return fmt.Errorf("can not write binary data to file : %w", err)

	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can not bianry data from file %w", err)
	}
	err = proto.Unmarshal(data, message)

	if err != nil {
		return fmt.Errorf("can not unmarshal binary to proto message %w", err)
	}

	return nil

}
