package auth

import (
	"context"

	"github.com/IbnAnjung/dealls/entity/enauth"
	"github.com/IbnAnjung/dealls/entity/enuser"
	coreerror "github.com/IbnAnjung/dealls/pkg/error"

	"github.com/IbnAnjung/dealls/pkg/jwt"
)

type RegisterInputValidation struct {
	Username        string `validate:"required,ascii,max=50,min=3"`
	Fullname        string `validate:"required,max=50,min=3"`
	Age             uint8  `validate:"required,max=60,min=16"`
	Gender          *uint8 `validate:"required,oneof=0 1"`
	Password        string `validate:"required,min=6,max=72"`
	ConfirmPassword string `validate:"required,min=6,eqfield=Password"`
}

func (uc *AuthUsecase) RegisterUser(ctx context.Context, input enauth.RegisterInput) (output enauth.RegisterOutput, err error) {
	// validate input
	if err = uc.validator.Validate(RegisterInputValidation{
		Fullname:        input.Fullname,
		Username:        input.Username,
		Age:             input.Age,
		Gender:          input.Gender,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	}); err != nil {
		return
	}

	// find user
	user, err := uc.userRepository.FindUserByUsername(ctx, input.Username)
	if err != nil {
		return
	}

	if user.ID != "" {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "username registered")
		err = e
		return
	}

	// create new user
	hashedPassword, err := uc.hasher.HashString(input.Password)
	if err != nil {
		return
	}

	user = enuser.User{
		ID:        uc.randomString.GenerateID(),
		Username:  input.Username,
		Fullname:  input.Fullname,
		Age:       input.Age,
		Gender:    enuser.UserGender(*input.Gender),
		Password:  hashedPassword,
		IsPremium: false,
		CreatedAt: uc.timeService.Now(),
	}

	i := 1

	for {
		if err = uc.userRepository.Create(ctx, &user); err != nil {
			if coreerror.GetType(err) != coreerror.CoreErrorDuplicate || i == 3 {
				err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "fail generate new user")
				return
			}
			user.ID = uc.randomString.GenerateID()
			i++
		} else {
			break
		}
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

	output = enauth.RegisterOutput{
		ID:           user.ID,
		Username:     user.Username,
		Fullname:     user.Fullname,
		Gender:       int8(user.Gender),
		Age:          user.Age,
		IsPremium:    user.IsPremium,
		CreatedAt:    user.CreatedAt,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
	return
}
