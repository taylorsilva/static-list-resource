package resource

type Request struct {
	Source  Source      `json:"source"`
	Version interface{} `json:"version"`
}

type Source struct {
	List []interface{} `json:"list"`
}
