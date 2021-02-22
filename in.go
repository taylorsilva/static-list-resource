package resource

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

func NewIn() in {
	return in{}
}

type in struct{}

func (i *in) Run(request InRequest, path string) (InResponse, error) {
	for _, item := range request.Source.List {
		if reflect.DeepEqual(item, request.Version.Item) {
			err := writeItemToFile(item, path)
			if err != nil {
				return InResponse{}, err
			}
			return InResponse{Version: Version{Item: item}}, nil
		}
	}
	return InResponse{}, errors.New("selected item not found in source.list")
}

func writeItemToFile(item interface{}, path string) error {
	filepath := filepath.Join(path, "item")
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	itm := []byte(fmt.Sprintf("%s", item))
	_, err = file.Write(itm)
	if err != nil {
		return err
	}

	return nil
}
