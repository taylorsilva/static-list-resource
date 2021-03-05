package resource

import (
	"errors"
	"time"
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
		return CheckResponse{
			Version{
				Item: request.Source.List[0],
				Date: time.Now(),
			},
		}, nil
	}

	for i, item := range request.Source.List {
		if item == request.Version.Item {
			if (i + 1) == len(request.Source.List) {
				// reached end of list, return first item
				return CheckResponse{
					Version{
						Item: request.Source.List[0],
						Date: time.Now(),
					},
				}, nil
			}
			return CheckResponse{
				Version{
					Item: request.Source.List[i+1],
					Date: time.Now(),
				},
			}, nil
		}
	}

	return CheckResponse{
		Version{
			Item: request.Source.List[0],
			Date: time.Now(),
		},
	}, nil
}

func convertList(list []string) CheckResponse {
	var response CheckResponse
	for _, v := range list {
		response = append([]Version{{Item: v}}, response...)
	}
	return response
}
