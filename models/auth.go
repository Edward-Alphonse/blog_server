package models

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id       int64  `db:"unique_id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
