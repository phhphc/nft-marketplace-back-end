package dto

type PagedRespond[T any] struct {
	PageSize    uint `json:"pageSize"`
	CurrentPage uint `json:"pageNum"`
	Content     []T  `json:"content"`
}
