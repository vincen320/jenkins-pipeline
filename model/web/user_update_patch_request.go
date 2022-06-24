package web

type UserUpdatePatchRequest struct {
	Id         int    `json:"id,omitempty"`
	Username   string `validate:"max=20,min=6" json:"username,omitempty"`
	Password   string `validate:"max=20,min=6" json:"password,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty"`
	LastOnline int64  `json:"lastOnline,omitempty"`
}
