package internal

import (
	"google.golang.org/protobuf/proto"
	"log"
	impl_v1 "protobuf-dynamic-go/impl.v1"
	"testing"
)

func TestMessageCast12(t *testing.T) {
	var msgCast1 = impl_v1.MsgCast1{
		Col1: 111,
		Col2: 222,
	}
	printObj(&msgCast1)

	binPack1, err := proto.Marshal(&msgCast1)
	if err != nil {
		log.Fatal(err)
	}

	msg2 := impl_v1.MsgCast2{}
	if err := proto.Unmarshal(binPack1, &msg2); err != nil {
		t.Fatal(err)
	}
	printObj(&msg2)

	msg3 := impl_v1.MsgCast3{}
	if err := proto.Unmarshal(binPack1, &msg3); err != nil {
		t.Fatal(err)
	}
	printObj(&msg3)
}

func TestMessageCast21(t *testing.T) {
	var msgCast2 = impl_v1.MsgCast2{
		Col1: 111,
		Col2: 222,
	}
	printObj(&msgCast2)

	binPack1, err := proto.Marshal(&msgCast2)
	if err != nil {
		log.Fatal(err)
	}

	msg1 := impl_v1.MsgCast1{}
	if err := proto.Unmarshal(binPack1, &msg1); err != nil {
		t.Fatal(err)
	}
	printObj(&msg1)

	msg3 := impl_v1.MsgCast3{}
	if err := proto.Unmarshal(binPack1, &msg3); err != nil {
		t.Fatal(err)
	}
	printObj(&msg3)
}
