package request

type CreateUser struct {
	Username string `json:"username" validate:"required,min=10,max=50"`
	Email    string `json:"email"  validate:"required,email,min=10,max=50"`
	Password string `json:"password" validate:"required,email,min=8,max=128"`
}

type UpdateUser struct {
	Username string `json:"username" validate:"required,min=10,max=50"`
	Password string `json:"password" validate:"required,email,min=8,max=128"`
}
