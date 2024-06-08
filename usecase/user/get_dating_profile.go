package user

import (
	"context"
	"fmt"

	"github.com/IbnAnjung/dealls/entity/enuser"
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

	// find recomendation_user
	recUser, err = uc.userRepository.FindUsers(ctx, uint8(user.Gender.GetOposite()), append([]string{user.ID}, swippedUserIDs...))
	if err != nil {
		return
	}

	return
}
