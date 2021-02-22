package resource

import (
	"errors"
)

func NewCheck() check {
	return check{}
}

type check struct{}

func (c *check) Run(request CheckRequest) (CheckResponse, error) {
	if len(request.Source.List) == 0 {
		return nil, errors.New("empty list provided in resouce's Source. At least one item required")
	}

	if request.Version.Item == "" {
		return convertList(request.Source.List), nil
	}

	// I wonder if there's a nice way to unmarshal an array
	// into a container/ring
	for i, item := range request.Source.List {
		if item == request.Version.Item {
			if (i + 1) == len(request.Source.List) {
				// reached end of list, return first item
				return convertList([]string{request.Source.List[0]}), nil
			}
			return convertList([]string{request.Source.List[i+1]}), nil
		}
	}

	return convertList([]string{request.Source.List[0]}), nil
}

func convertList(list []string) CheckResponse {
	var response CheckResponse
	for _, v := range list {
		response = append([]Version{{Item: v}}, response...)
	}
	return response
}
