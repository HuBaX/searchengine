package util

import (
	"encoding/json"
	"io"
	"log"
)

func JsonEncode(source interface{}, output io.Writer) {
	if err := json.NewEncoder(output).Encode(source); err != nil {
		log.Printf("Error encoding %v to json: %s", source, err)
	}
}

func JsonDecode(source io.ReadCloser, output interface{}) {
	if err := json.NewDecoder(source).Decode(output); err != nil {
		log.Printf("Error parsing %v to interface{}: %s", source, err)
	}
}
