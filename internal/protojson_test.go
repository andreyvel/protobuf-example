package internal

import (
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	impl_v1 "protobuf-dynamic-go/impl.v1"
	"testing"
)

func TestProtojson(t *testing.T) {
	var msgV1 = getMsgV1()
	data, err := protojson.Marshal(msgV1)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr := string(data)
	log.Println(jsonStr)

	{
		msgOut := impl_v1.Message{}
		err = protojson.Unmarshal([]byte(jsonStr), &msgOut)
		if err != nil {
			log.Fatal(err)
		}
		printObj(msgOut)
	}

	{
		msgOut := impl_v1.Message{}
		jsonStr = `{"id":7777,"name":"John Doe","email":"jdoe@example.com","phones":[{"number":"333-12345"}]}`
		err = protojson.Unmarshal([]byte(jsonStr), &msgOut)
		if err != nil {
			log.Fatal(err)
		}
		printObj(msgOut)
	}
}
