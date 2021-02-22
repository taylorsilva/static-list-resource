package resource

type Source struct {
	List []interface{} `json:"list"`
}

type Version struct {
	Item interface{} `json:"item"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type CheckResponse []Version

type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InResponse struct {
	Version Version `json:"version"`
}
