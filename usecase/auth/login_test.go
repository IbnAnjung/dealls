package auth

import (
	"context"
	"testing"
	"time"

	"github.com/IbnAnjung/dealls/entity/enauth"
	entityUser "github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/entity/enuser"
	"github.com/IbnAnjung/dealls/mock/pkg/crypt"
	"github.com/IbnAnjung/dealls/mock/pkg/jwt"
	pkgString "github.com/IbnAnjung/dealls/mock/pkg/string"
	"github.com/IbnAnjung/dealls/mock/pkg/structvalidator"
	pkgTime "github.com/IbnAnjung/dealls/mock/pkg/time"
	pkgJwt "github.com/IbnAnjung/dealls/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	hasher := crypt.NewMockHash(t)
	jwtServ := jwt.NewMockJwtService(t)
	validator := structvalidator.NewMockValidator(t)
	timeServ := pkgTime.NewMockTime(t)
	randString := pkgString.NewMockRandomString(t)
	userRepo := enuser.NewMockUserRepository(t)

	input := enauth.LoginInput{
		Username: "dealls",
		Password: "password",
	}

	id := "user-rand-uuid"
	cAt := time.Now()
	hashedPassword := "hashed-password"
	accToken := "access-token"
	refhToken := "refresh-token"

	user := entityUser.User{
		ID:        id,
		Fullname:  "Dealls - sejuta cita",
		Username:  input.Username,
		Age:       20,
		Gender:    entityUser.UserGenderMale,
		Password:  hashedPassword,
		IsPremium: false,
		CreatedAt: cAt,
	}

	uClaim := pkgJwt.UserClaim{
		UserID:    user.ID,
		IsPremium: user.IsPremium,
	}

	validator.EXPECT().Validate(LoginInputValidation{
		Username: input.Username,
		Password: input.Password,
	}).Return(nil)
	userRepo.EXPECT().FindUserByUsername(context.Background(), input.Username).Return(user, nil)
	hasher.EXPECT().CompareHash(hashedPassword, input.Password).Return(nil)
	jwtServ.EXPECT().GenerateAccessToken(uClaim).Return(accToken, nil)
	jwtServ.EXPECT().GenerateRefreshToken(uClaim).Return(refhToken, nil)

	uc := NewUsecase(hasher, jwtServ, validator, timeServ, randString, userRepo)

	output, err := uc.Login(context.Background(), input)

	assert.Nil(t, err)
	assert.Equal(t, enauth.LoginOutput{
		ID:           user.ID,
		Username:     user.Username,
		Fullname:     user.Fullname,
		Gender:       output.Gender,
		Age:          user.Age,
		IsPremium:    user.IsPremium,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accToken,
		RefreshToken: refhToken,
	}, output)
}
