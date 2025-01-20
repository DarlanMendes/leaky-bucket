package model

type Pagination struct {
	Content interface{} `json:"content"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int64       `json:"total"`
}
