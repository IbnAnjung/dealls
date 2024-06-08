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

func TestRegister(t *testing.T) {
	hasher := crypt.NewMockHash(t)
	jwtServ := jwt.NewMockJwtService(t)
	validator := structvalidator.NewMockValidator(t)
	timeServ := pkgTime.NewMockTime(t)
	randString := pkgString.NewMockRandomString(t)
	userRepo := enuser.NewMockUserRepository(t)
	gender := uint8(entityUser.UserGenderMale)

	input := enauth.RegisterInput{
		Fullname:        "Dealls - sejuta cita",
		Username:        "dealls",
		Age:             20,
		Gender:          &gender,
		Password:        "password",
		ConfirmPassword: "password",
	}

	id := "user-rand-uuid"
	cAt := time.Now()
	hashedPassword := "hashed-password"
	accToken := "access-token"
	refhToken := "refresh-token"

	expectedUser := entityUser.User{
		ID:        id,
		Fullname:  input.Fullname,
		Username:  input.Username,
		Age:       input.Age,
		Gender:    entityUser.UserGender(*input.Gender),
		Password:  hashedPassword,
		IsPremium: false,
		CreatedAt: cAt,
	}

	uClaim := pkgJwt.UserClaim{
		UserID:    expectedUser.ID,
		IsPremium: expectedUser.IsPremium,
	}

	validator.EXPECT().Validate(RegisterInputValidation{
		Fullname:        input.Fullname,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
		Age:             input.Age,
		Gender:          input.Gender,
	}).Return(nil)
	userRepo.EXPECT().FindUserByUsername(context.Background(), input.Username).Return(entityUser.User{}, nil)
	hasher.EXPECT().HashString(input.Password).Return(hashedPassword, nil)
	randString.EXPECT().GenerateID().Return(id)
	timeServ.EXPECT().Now().Return(cAt)
	userRepo.EXPECT().Create(context.Background(), &expectedUser).Return(nil)
	jwtServ.EXPECT().GenerateAccessToken(uClaim).Return(accToken, nil)
	jwtServ.EXPECT().GenerateRefreshToken(uClaim).Return(refhToken, nil)

	uc := NewUsecase(hasher, jwtServ, validator, timeServ, randString, userRepo)

	output, err := uc.RegisterUser(context.Background(), input)

	assert.Nil(t, err)
	assert.Equal(t, enauth.RegisterOutput{
		ID:           expectedUser.ID,
		Username:     expectedUser.Username,
		Fullname:     expectedUser.Fullname,
		Gender:       output.Gender,
		Age:          expectedUser.Age,
		IsPremium:    expectedUser.IsPremium,
		CreatedAt:    expectedUser.CreatedAt,
		AccessToken:  accToken,
		RefreshToken: refhToken,
	}, output)
}
