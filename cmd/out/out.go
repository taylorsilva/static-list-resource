package main

import (
	"encoding/json"
	"fmt"
	"os"

	resource "github.com/taylorsilva/static-list-resource"
)

func main() {
	var request resource.OutRequest
	decorder := json.NewDecoder(os.Stdin)
	decorder.DisallowUnknownFields()
	err := decorder.Decode(&request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal put request: %s", err)
		os.Exit(1)
	}

	out := resource.NewOut()
	//TODO: error and nil checking
	path := os.Args[1]
	response, err := out.Run(request, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "put error: %s", err)
		os.Exit(1)
	}

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding response to json: %s", err)
		os.Exit(1)
	}
}
