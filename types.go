package resource

import "time"

type Source struct {
	List []string `json:"list"`
}

type Version struct {
	Item string    `json:"item"`
	Date time.Time `json:"date"`
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
