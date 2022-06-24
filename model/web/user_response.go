package web

type UserResponse struct {
	Id         int    `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	CreatedAt  int64  `json:"createdAt,omitempty"`
	LastOnline int64  `json:"lastOnline,omitempty"`
}
