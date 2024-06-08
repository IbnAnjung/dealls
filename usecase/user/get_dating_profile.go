package user

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/dealls/entity/enuser"
	coreerror "github.com/IbnAnjung/dealls/pkg/error"
)

func (uc *userUsecase) GetDatingUserProfile(ctx context.Context, userID string) (recUser enuser.User, err error) {
	swippedUserIDs := []string{}
	if err = uc.cacheService.Get(ctx, fmt.Sprintf("swipped_user_%s", userID), &swippedUserIDs); err != nil {
		fmt.Println("error get val", err.Error())
		return
	}

	user, err := uc.userRepository.FindById(ctx, userID)
	if err != nil {
		return
	}

	// return err if > 10 and isPremium false
	if len(swippedUserIDs) > 10 && !user.IsPremium {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "swipe exceeds daily limit")
		return
	}

	// find recomendation_user
	recUser, err = uc.userRepository.FindUsers(ctx, uint8(user.Gender.GetOposite()), append([]string{user.ID}, swippedUserIDs...))
	if err != nil {
		return
	}

	return
}
