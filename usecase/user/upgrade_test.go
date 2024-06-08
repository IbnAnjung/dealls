package user

import (
	"context"
	"testing"

	entityUser "github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/pkg/cache"
	"github.com/IbnAnjung/dealls/mock/pkg/structvalidator"
	pkgTime "github.com/IbnAnjung/dealls/mock/pkg/time"
	"github.com/stretchr/testify/assert"
)

func TestUpgrade(t *testing.T) {
	validator := structvalidator.NewMockValidator(t)
	timeServ := pkgTime.NewMockTime(t)
	cacheServ := cache.NewMockCacheService(t)
	userRepo := enuser.NewMockUserRepository(t)

	ctx := context.Background()
	user := entityUser.User{
		ID:        "user-id",
		Gender:    entityUser.UserGenderFemale,
		IsPremium: false,
	}

	userRepo.EXPECT().FindById(ctx, user.ID).Return(user, nil)
	user.IsPremium = true
	userRepo.EXPECT().Update(ctx, &user).Return(nil)

	uc := NewUsecase(validator, timeServ, cacheServ, userRepo)

	err := uc.UpgradeToPremium(ctx, user.ID)
	assert.Nil(t, err)
}
