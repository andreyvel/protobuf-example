package internal

import (
	"google.golang.org/protobuf/proto"
	"log"
	impl_v1 "protobuf-dynamic-go/impl.v1"
	"testing"
)

func TestMessagePack(t *testing.T) {
	var msgPack1 = impl_v1.MsgPack1{
		Msg: &impl_v1.MessageInner{
			Id:    3333,
			Name:  "John Doe",
			Email: "jdoe@example.com",
		},
		Id: "333",
	}

	binPack1, err := proto.Marshal(&msgPack1)
	if err != nil {
		log.Fatal(err)
	}

	msgV1 := impl_v1.MsgPack1{}
	if err := proto.Unmarshal(binPack1, &msgV1); err != nil {
		t.Fatal(err)
	}
	printObj(&msgV1)
	// {"msg":{"id":3333,"name":"John Doe","email":"jdoe@example.com"},"id":"333"}

	msgV2 := impl_v1.MsgPack2{}
	if err := proto.Unmarshal(binPack1, &msgV2); err != nil {
		t.Fatal(err)
	}
	printObj(&msgV2)
	// {"msg":"CIUaEghKb2huIERvZRoQamRvZUBleGFtcGxlLmNvbQ==","id":"333"}

	msgV3 := impl_v1.MsgPack3{}
	if err := proto.Unmarshal(binPack1, &msgV3); err != nil {
		t.Fatal(err)
		// msgpack _test.go:39: string field contains invalid UTF-8
	}
	printObj(&msgV3)
}
