package model

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Token string `db:"token" json:"token"`
}

type UserCreateRequest struct {
	Name string `db:"name" json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Name string `db:"name" json:"name"`
}

type UserUpdateRequest struct {
	Name string `db:"name" json:"name"`
}
