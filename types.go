package resource

import "time"

type Source struct {
	List []string `json:"list"`
	// A hack to avoid global resources behaviour
	Unique string `json:"unique"`
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

type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

type OutParams struct {
	// Path to a file containing the previous version
	Previous string `json:"previous"`
}

type OutResponse struct {
	Version Version `json:"version"`
}
