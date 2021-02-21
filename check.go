package resource

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func NewCheck() check {
	return check{}
}

type check struct{}

func (c *check) Run(request Request) ([]interface{}, error) {
	if len(request.Source.List) == 0 {
		return nil, errors.New("empty list provided. At least one item required")
	}

	if request.Version == nil {
		return request.Source.List, nil
	}

	// I wonder if there's a nice way to unmarshal an array
	// into a container/ring
	for i, item := range request.Source.List {
		if item == request.Version {
			if (i + 1) == len(request.Source.List) {
				// reached end of list, return first item
				return []interface{}{request.Source.List[0]}, nil
			}
			return []interface{}{request.Source.List[i+1]}, nil
		}
	}

	return []interface{}{request.Source.List[0]}, nil
}
