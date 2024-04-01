package internal

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	impl_v1 "protobuf-dynamic-go/impl.v1"
	impl_v2 "protobuf-dynamic-go/impl.v2"
	"testing"
)

func TestDecoder(t *testing.T) {
	var msgV1 = getMsgV1()
	dataV1, err := proto.Marshal(msgV1)
	if err != nil {
		log.Fatal(err)
	}

	var msgV2 = getMsgV2()
	dataV2, err := proto.Marshal(msgV2)
	if err != nil {
		log.Fatal(err)
	}

	decodeV1(dataV1)
	decodeV1(dataV2)

	decodeV2(dataV1)
	decodeV2(dataV2)
}

func TestMarshalToFile(t *testing.T) {
	var msgV1 = getMsgV1()

	dataOut, err := proto.Marshal(msgV1)
	if err != nil {
		log.Fatal(err)
	}

	fileOut, err := os.CreateTemp(os.TempDir(), "*.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer fileOut.Close()

	_, err = fileOut.Write(dataOut)
	if err != nil {
		log.Fatalln(err)
	}

	fileName := fileOut.Name()
	dataIn, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	msgOut := impl_v1.Message{}
	if err := proto.Unmarshal(dataIn, &msgOut); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, msgV1.String(), msgOut.String())
}

func getMsgV1() *impl_v1.Message {
	var msg = impl_v1.Message{
		Id:          3333,
		Name:        "John Doe",
		Email:       "jdoe@example.com",
		Phones:      []*impl_v1.Message_PhoneNumber{{Type: impl_v1.Message_MOBILE, Number: "333-12345"}},
		LastUpdated: timestamppb.Now(),
	}
	return &msg
}

func getMsgV2() *impl_v2.Message {
	var msg = impl_v2.Message{
		Id:          5555,
		Name:        "John Doe",
		Email:       "jdoe@example.com",
		Phones:      []*impl_v2.Message_PhoneNumber{{Type: impl_v2.Message_MOBILE, Number: "555-12345"}},
		LastUpdated: timestamppb.Now(),
		Desc:        "desc",
	}
	return &msg
}

func decodeV1(data []byte) {
	msgV1 := impl_v1.Message{}
	if err := proto.Unmarshal(data, &msgV1); err != nil {
		log.Fatalln(err)
	}
	printObj(&msgV1)
}

func decodeV2(data []byte) {
	msg := impl_v2.Message{}
	if err := proto.Unmarshal(data, &msg); err != nil {
		log.Fatalln(err)
	}
	printObj(&msg)
}

func printObj(msg any) {
	jsonV1, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	println(string(jsonV1))
}
