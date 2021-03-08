package resource

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

func NewOut() out {
	return out{}
}

type out struct{}

func (c *out) Run(request OutRequest, artifactsPath string) (OutResponse, error) {
	if len(request.Source.List) == 0 {
		return OutResponse{}, errors.New("empty list provided in resouce's Source. At least one item required")
	}
	if request.Params.Previous == "" {
		return OutResponse{}, errors.New("no file provided with previous version")
	}

	prevFile := filepath.Join(artifactsPath, request.Params.Previous)
	contents, err := os.ReadFile(prevFile)
	if err != nil {
		return OutResponse{}, err
	}

	prevVersion := string(contents)
	var nextVersion string

	if prevVersion == "" {
		nextVersion = request.Source.List[0]
	}

	for i, v := range request.Source.List {
		if v == prevVersion {
			if i == len(request.Source.List)-1 {
				nextVersion = request.Source.List[0]
				break
			}
			nextVersion = request.Source.List[i+1]
			break
		}
	}

	return OutResponse{
		Version: Version{
			Item: nextVersion,
			Date: time.Now(),
		},
	}, nil
}
