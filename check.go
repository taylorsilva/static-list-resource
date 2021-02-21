package resource

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func NewCheck() check {
	return check{}
}

type check struct{}

func (c *check) Run(request Request) []interface{} {
	if request.Version == nil {
		return request.Source.List
	}
	// I wonder if there's a nice way to unmarshal an array
	// into a container/ring
	for i, item := range request.Source.List {
		if item == request.Version {
			if (i + 1) == len(request.Source.List) {
				return []interface{}{request.Source.List[0]}
			}
			return []interface{}{request.Source.List[i+1]}
		}
	}
	return []interface{}{}
}
