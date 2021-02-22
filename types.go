package resource

type CheckRequest struct {
	Source  Source      `json:"source"`
	Version interface{} `json:"version"`
}

type Source struct {
	List []interface{} `json:"list"`
}
