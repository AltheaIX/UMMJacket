package model

type Statistic struct {
	UsersCount        int `db:"users_count" json:"usersCount"`
	TransactionsCount int `db:"transactions_count" json:"transactionsCount"`
}
