package web

type UserCreateRequest struct {
	Username string `validate:"required,max=20,min=6" json:"username"`
	Password string `validate:"required,max=20,min=6" json:"password"`
}
