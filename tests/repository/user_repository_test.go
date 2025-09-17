package repository

import (
	"context"
	"github.com/AltheaIX/UMMJacket/configs"
	"github.com/AltheaIX/UMMJacket/infras"
	"github.com/AltheaIX/UMMJacket/internal/domain/user/repository"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"testing"
)

func TestGetUsers(t *testing.T) {
	cfg := configs.GetConfig()

	db, _ := infras.InitMysql(cfg)
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	//groupFilter := []filter.Filter{
	//	{
	//		Field:    "nim",
	//		Operator: "eq",
	//		Value:    "202410370110031",
	//	},
	//}
	//
	//filters := filter.FiltersJacket{Filter: groupFilter}

	data, err := userRepo.ResolveUsersRepository(context.Background(), &filter.Filters{})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
	return
}
