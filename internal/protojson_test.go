package internal

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	impl_v1 "protobuf-dynamic-go/impl.v1"
	"testing"
)

func TestProtojson(t *testing.T) {
	var msg = &impl_v1.Message{
		Id:          3333,
		Name:        "John Doe",
		Email:       "jdoe@example.com",
		LastUpdated: timestamppb.Now(),
	}

	data, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr := string(data)
	log.Println("\n" + jsonStr)
	//{"id":3333, "name":"John Doe", "email":"jdoe@example.com", "lastUpdated":"2024-04-20T07:25:06.644458753Z"}

	{
		msgOut := impl_v1.Message{}
		err = protojson.Unmarshal([]byte(jsonStr), &msgOut)
		if err != nil {
			log.Fatal(err)
		}

		printObj(&msgOut)
		// Field lastUpdated converted to last_updated structure
		//{"id":3333,"name":"John Doe","email":"jdoe@example.com","last_updated":{"seconds":1713597906,"nanos":644458753}}
	}

	{
		msgOut := impl_v1.Message{}
		jsonStr = `{"id":7777,"name":"John Doe","email":"jdoe@example.com"}`
		err = protojson.Unmarshal([]byte(jsonStr), &msgOut)
		if err != nil {
			log.Fatal(err)
		}

		printObj(&msgOut)
		//{"id":7777,"name":"John Doe","email":"jdoe@example.com"}
	}
}
