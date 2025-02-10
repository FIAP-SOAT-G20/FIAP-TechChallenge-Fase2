package usecase

type PaginatedOutput struct {
	Total int64
	Page  int
	Limit int
}
