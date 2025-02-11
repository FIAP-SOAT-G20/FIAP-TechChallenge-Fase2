package dto

type JsonPagination struct {
	Total int64 `json:"total" example:"100"`
	Page  int   `json:"page" example:"1"`
	Limit int   `json:"limit" example:"10"`
}

type XmlPagination struct {
	Total int64 `xml:"total" example:"100"`
	Page  int   `xml:"page" example:"1"`
	Limit int   `xml:"limit" example:"10"`
}
