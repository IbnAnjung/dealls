package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	entityUser "github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/pkg/cache"
	"github.com/IbnAnjung/dealls/mock/pkg/structvalidator"
	pkgTime "github.com/IbnAnjung/dealls/mock/pkg/time"
	"github.com/stretchr/testify/assert"
)

func TestSwipe(t *testing.T) {
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

	swapUser := entityUser.User{
		ID:        "swao-user-id",
		Gender:    entityUser.UserGenderMale,
		IsPremium: false,
	}

	swType := entityUser.SwipeTypeLike
	input := entityUser.SwipeUserInput{
		UserID:        user.ID,
		SwippedUserID: swapUser.ID,
		Type:          &swType,
	}

	now := time.Now()
	lastDay := time.Date(now.Year(), now.Month(), now.Day(), 24, 59, 59, 59, time.Local)
	duration := lastDay.Sub(now)

	validator.EXPECT().Validate(SwipeInputValidation{
		UserID:       input.UserID,
		SwippedUseID: swapUser.ID,
		Type:         input.Type,
	}).Return(nil)
	cacheServ.EXPECT().GetAndLock(ctx, fmt.Sprintf("swipped_user_%s", user.ID), &cacheValuses).Return(nil)
	userRepo.EXPECT().FindByIds(ctx, []string{user.ID, swapUser.ID}).Return([]entityUser.User{user, swapUser}, nil)
	userRepo.EXPECT().UpdateLikeCount(ctx, swapUser.ID, *input.Type).Return(nil)
	timeServ.EXPECT().Now().Return(now)
	timeServ.EXPECT().GetLastTimeOnDay().Return(lastDay)
	cacheServ.EXPECT().Set(ctx, fmt.Sprintf("swipped_user_%s", input.UserID), append(cacheValuses, swapUser.ID), duration).Return(nil)
	cacheServ.EXPECT().Unlock(ctx, fmt.Sprintf("swipped_user_%s", input.UserID)).Return(nil)

	uc := NewUsecase(validator, timeServ, cacheServ, userRepo)

	err := uc.Swipe(ctx, input)
	assert.Nil(t, err)
}
