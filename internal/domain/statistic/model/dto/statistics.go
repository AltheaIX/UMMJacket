package dto

type DashboardResponse struct {
	UsersCount        int `json:"usersCount"`
	TransactionsCount int `json:"transactionsCount"`
}
