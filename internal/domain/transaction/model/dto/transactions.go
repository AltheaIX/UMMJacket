package dto

type NewTransactionRequest struct {
	ProductId int `json:"productId"`
}

type NewTransactionResponse struct {
	Id        int64 `db:"id"  json:"id"`
	UserId    int   `json:"userId"`
	ProductId int   `json:"productId"`
}
