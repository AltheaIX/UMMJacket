package repository

import (
	"context"
	"github.com/AltheaIX/UMMJacket/infras"
	"github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"testing"
)

func TestGetUsers(t *testing.T) {
	db, _ := infras.InitPostgres()
	defer db.Close()

	userRepo := repository.NewUserRepositoryImpl(db)

	//groupFilter := []filter.Filter{
	//	{
	//		Field:    "nim",
	//		Operator: "eq",
	//		Value:    "202410370110031",
	//	},
	//}
	//
	//filters := filter.Filters{Filter: groupFilter}

	data, err := userRepo.ResolveUsersRepository(context.Background(), &filter.Filters{})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
	return
}
