package model

type HardDisk struct {
	Name  string `json:"name"`
	Total string `json:"total"`
	Used  string `json:"used"`
	Free  string `json:"free"`
}
