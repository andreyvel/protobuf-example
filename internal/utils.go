package internal

import (
	"encoding/json"
	"log"
)

func printObj(msg any) {
	jsonV1, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	println(string(jsonV1))
}
