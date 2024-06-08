package user

import (
	"context"
	"fmt"
	"testing"

	entityUser "github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/pkg/cache"
	"github.com/IbnAnjung/dealls/mock/pkg/structvalidator"
	pkgTime "github.com/IbnAnjung/dealls/mock/pkg/time"
	"github.com/stretchr/testify/assert"
)

func TestGetDatingProfile(t *testing.T) {
	validator := structvalidator.NewMockValidator(t)
	timeServ := pkgTime.NewMockTime(t)
	cacheServ := cache.NewMockCacheService(t)
	userRepo := enuser.NewMockUserRepository(t)

	ctx := context.Background()
	cacheValuses := []string{}
	user := entityUser.User{
		ID:        "user-id",
		Gender:    entityUser.UserGenderFemale,
		IsPremium: false,
	}

	expectedUser := entityUser.User{
		ID:        "date-user-id",
		Gender:    entityUser.UserGenderMale,
		IsPremium: false,
	}

	cacheServ.EXPECT().Get(ctx, fmt.Sprintf("swipped_user_%s", user.ID), &cacheValuses).Return(nil)
	userRepo.EXPECT().FindById(ctx, user.ID).Return(user, nil)
	userRepo.EXPECT().FindUsers(ctx, uint8(user.Gender.GetOposite()), append([]string{user.ID}, cacheValuses...)).Return(expectedUser, nil)

	uc := NewUsecase(validator, timeServ, cacheServ, userRepo)

	output, err := uc.GetDatingUserProfile(ctx, user.ID)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, output)
}
