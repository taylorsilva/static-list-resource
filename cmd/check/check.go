package main

import (
	"encoding/json"
	"fmt"
	"os"

	resource "github.com/taylorsilva/static-list-resource"
)

func main() {
	var request resource.CheckRequest
	decorder := json.NewDecoder(os.Stdin)
	decorder.DisallowUnknownFields()
	err := decorder.Decode(&request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal check request: %s", err)
		os.Exit(1)
	}

	check := resource.NewCheck()
	response, err := check.Run(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check error: %s", err)
		os.Exit(1)
	}

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(response)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding response to json: %s", err)
		os.Exit(1)
	}
}
