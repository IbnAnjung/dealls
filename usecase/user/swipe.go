package user

import (
	"context"
	"fmt"
	"slices"

	"github.com/IbnAnjung/dealls/entity/enuser"
	coreerror "github.com/IbnAnjung/dealls/pkg/error"
)

type SwipeInputValidation struct {
	UserID       string            `validate:"required"`
	SwippedUseID string            `validate:"required"`
	Type         *enuser.SwipeType `validate:"required,oneof=0 1"`
}

func (uc *userUsecase) Swipe(ctx context.Context, input enuser.SwipeUserInput) (err error) {
	if err = uc.validator.Validate(SwipeInputValidation{
		UserID:       input.UserID,
		SwippedUseID: input.SwippedUserID,
		Type:         input.Type,
	}); err != nil {
		return
	}

	// get swipped user from cache
	swippedUserIDs := []string{}
	if err = uc.cacheService.GetAndLock(ctx, fmt.Sprintf("swipped_user_%s", input.UserID), &swippedUserIDs); err != nil {
		return
	}

	defer func() {
		if e := uc.cacheService.Unlock(ctx, fmt.Sprintf("swipped_user_%s", input.UserID)); e != nil {
			fmt.Printf("fail unlock user profile: %s \n", e.Error())
			err = e
		}
	}()

	if slices.Contains(swippedUserIDs, input.SwippedUserID) {
		return coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "profile already swipe")
	}

	var user, swUser enuser.User
	users, err := uc.userRepository.FindByIds(ctx, []string{input.UserID, input.SwippedUserID})
	if err != nil {
		return
	}

	for _, v := range users {
		if v.ID == input.UserID {
			user = v
		} else {
			swUser = v
		}
	}

	if user.ID == "" || swUser.ID == "" {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "invalid profile")
		return
	}

	// return err if > 10 and isPremium false
	if len(swippedUserIDs) >= 10 && !user.IsPremium {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "swipe exceeds daily limit")
		return
	}

	if err = uc.userRepository.UpdateLikeCount(ctx, input.SwippedUserID, *input.Type); err != nil {
		return
	}

	swippedUserIDs = append(swippedUserIDs, swUser.ID)

	if err = uc.cacheService.Set(ctx, fmt.Sprintf("swipped_user_%s", input.UserID), swippedUserIDs, uc.timeService.GetLastTimeOnDay().Sub(uc.timeService.Now())); err != nil {
		return
	}

	return
}
