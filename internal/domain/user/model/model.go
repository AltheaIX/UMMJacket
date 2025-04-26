package model

type User struct {
	Id        int    `db:"id"  json:"id"`
	Nim       string `db:"nim" json:"nim"`
	Password  string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}
