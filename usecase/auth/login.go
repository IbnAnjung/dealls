package auth

import (
	"context"

	"github.com/IbnAnjung/dealls/entity/enauth"
	coreerror "github.com/IbnAnjung/dealls/pkg/error"

	"github.com/IbnAnjung/dealls/pkg/jwt"
)

type LoginInputValidation struct {
	Username string `validate:"required,ascii,max=50"`
	Password string `validate:"required"`
}

func (uc *AuthUsecase) Login(ctx context.Context, input enauth.LoginInput) (output enauth.LoginOutput, err error) {
	// validate input
	if err = uc.validator.Validate(LoginInputValidation{
		Username: input.Username,
		Password: input.Password,
	}); err != nil {
		return
	}

	// find user
	user, err := uc.userRepository.FindUserByUsername(ctx, input.Username)
	if err != nil {
		return
	}

	if user.ID == "" {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "invalid credentials")
		err = e
		return
	}

	// create new user
	if err = uc.hasher.CompareHash(user.Password, input.Password); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "invalid credentials")
		err = e
		return
	}

	// generate token
	userClaim := jwt.UserClaim{
		UserID:    user.ID,
		IsPremium: user.IsPremium,
	}

	accessToken, err := uc.jwtService.GenerateAccessToken(userClaim)
	if err != nil {
		return
	}

	refreshToken, err := uc.jwtService.GenerateRefreshToken(userClaim)
	if err != nil {
		return
	}

	output = enauth.LoginOutput{
		ID:           user.ID,
		Username:     user.Username,
		Fullname:     user.Fullname,
		Gender:       int8(user.Gender),
		Age:          user.Age,
		IsPremium:    user.IsPremium,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return
}
