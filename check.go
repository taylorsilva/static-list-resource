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

	for _, item := range request.Source.List {
		if item == request.Version.Item {
			return CheckResponse{}, nil
		}
	}

	return CheckResponse{
		Version{
			Item: request.Source.List[0],
			Date: time.Now(),
		},
	}, nil
}
