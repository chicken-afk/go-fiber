package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
