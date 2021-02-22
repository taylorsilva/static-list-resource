package resource

func NewIn() in {
	return in{}
}

type in struct{}

func (i *in) Run(request InRequest) (InResponse, error) {
	return InResponse{Version: Version{Item: "item4"}}, nil
}
