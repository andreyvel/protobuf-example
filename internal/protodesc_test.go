package internal

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"log"
	"os"
	"os/exec"
	"path"
	"testing"
)

var schemaPath = "../schemas"

func TestProtoDesc(t *testing.T) {
	var msgV1 = getMsgV1()

	dataBin, err := proto.Marshal(msgV1)
	if err != nil {
		log.Fatal(err)
	}

	printObj(msgV1)
	decodeAndPrint(dataBin, "Message", schemaPath, "msg-v1.proto")
	decodeAndPrint(dataBin, "Message", schemaPath, "msg-v2.proto")

	var msgV2 = getMsgV2()
	dataBin, err = proto.Marshal(msgV2)
	if err != nil {
		log.Fatal(err)
	}

	printObj(msgV2)
	decodeAndPrint(dataBin, "Message", schemaPath, "msg-v1.proto")
	decodeAndPrint(dataBin, "Message", schemaPath, "msg-v2.proto")
}

func TestReflection(t *testing.T) {
	protoFile := "msg-v1.proto"
	registry, err := createProtoRegistry(schemaPath, protoFile)
	if err != nil {
		log.Fatal(err)
	}

	desc, err := registry.FindFileByPath(protoFile)
	if err != nil {
		log.Fatal(err)
	}

	msgs := desc.Messages()
	for ind := 0; ind < msgs.Len(); ind++ {
		msgDesc := msgs.Get(ind)
		println("--------------------------------------------------------------")
		fmt.Printf("-- %v, %v\n", msgDesc.Name(), msgDesc.FullName())
		println("--------------------------------------------------------------")

		fieldDesc := msgDesc.Fields()
		for ind := 0; ind < fieldDesc.Len(); ind++ {
			field := fieldDesc.Get(ind)
			fmt.Printf("%v, %v, %v\n", field.Name(), field.TextName(), field.FullName())
		}
	}
}

func decodeAndPrint(dataBin []byte, typeName string, schemaPath string, protoFile string) {
	registry, err := createProtoRegistry(schemaPath, protoFile)
	if err != nil {
		log.Fatal(err)
	}

	desc, err := registry.FindFileByPath(protoFile)
	if err != nil {
		log.Fatal(err)
	}

	fd := desc.Messages()
	messageDesc := fd.ByName(protoreflect.Name(typeName))

	if messageDesc == nil {
		log.Fatal("messageDesc == nil")
	}

	msg := dynamicpb.NewMessage(messageDesc)
	err = proto.Unmarshal(dataBin, msg)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}
}

func createProtoRegistry(srcDir string, filename string) (*protoregistry.Files, error) {
	tmpFile := filename + "-tmp.pb"
	cmd := exec.Command("protoc",
		"--include_imports",
		"--descriptor_set_out="+tmpFile,
		"-I"+srcDir,
		path.Join(srcDir, filename))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile)

	marshalledDescriptorSet, err := os.ReadFile(tmpFile)
	if err != nil {
		return nil, err
	}

	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
	if err != nil {
		return nil, err
	}

	files, err := protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}

	return files, nil
}
