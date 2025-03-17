package request

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Foods"`
}

type GetCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

type ListCategoriesRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,required" example:"Beverages"`
}

type DeleteCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}
