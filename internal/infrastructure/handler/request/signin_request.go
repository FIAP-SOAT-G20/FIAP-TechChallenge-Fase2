package request

type SignInRequest struct {
	CPF string `json:"cpf" binding:"required" example:"123.456.789-00"`
}
