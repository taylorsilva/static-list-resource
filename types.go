package resource

type Version struct {
	Item interface{} `json:"item"`
}

type CheckRequest struct {
	Source  Source      `json:"source"`
	Version interface{} `json:"version"`
}

type CheckResponse []Version

type Source struct {
	List []interface{} `json:"list"`
}
