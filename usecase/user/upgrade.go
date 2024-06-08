package user

import (
	"context"

	coreerror "github.com/IbnAnjung/dealls/pkg/error"
)

func (uc *userUsecase) UpgradeToPremium(ctx context.Context, userID string) (err error) {
	user, err := uc.userRepository.FindById(ctx, userID)
	if err != nil {
		return
	}

	if user.ID == "" {
		return coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "user not found")
	}

	if user.IsPremium {
		return coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "user already premium")
	}

	user.IsPremium = true

	if err = uc.userRepository.Update(ctx, &user); err != nil {
		return
	}

	return
}
