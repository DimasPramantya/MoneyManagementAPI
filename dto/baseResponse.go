package dto

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationResponse[T any] struct {
	TotalRecords  int    `json:"totalRecords"`
	TotalPages    int    `json:"totalPages"`
	CurrentPage   int    `json:"currentPage"`
	Limit         int    `json:"limit"`
	Records       []T    `json:"records"`
	NextPage      *int   `json:"nextPage"`
	PreviousPage  *int   `json:"previousPage"`
}