package model

type Transaction struct {
	Id        int    `db:"id" json:"id"`
	UserId    int    `db:"user_id" json:"userId"`
	ProductId int    `db:"product_id" json:"productId"`
	CreatedAt string `db:"created_at" json:"createdAt"`
	UpdatedAt string `db:"updated_at" json:"updatedAt"`
}
