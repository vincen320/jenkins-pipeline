package web

type UserUpdateRequest struct {
	Id         int    `json:"id,omitempty"`
	Username   string `validate:"required,omitempty,max=20,min=6" json:"username,omitempty"`
	Password   string `validate:"required,omitempty,max=20,min=6" json:"password,omitempty"`
	LastOnline int64  `validate:"required,omitempty" json:"lastOnline,omitempty"`
}
