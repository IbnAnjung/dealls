package auth

import (
	"github.com/IbnAnjung/dealls/entity/enauth"
	"github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/pkg/crypt"
	"github.com/IbnAnjung/dealls/pkg/string"
	"github.com/IbnAnjung/dealls/pkg/structvalidator"
	"github.com/IbnAnjung/dealls/pkg/time"

	"github.com/IbnAnjung/dealls/pkg/jwt"
)

type AuthUsecase struct {
	hasher         crypt.Hash
	jwtService     jwt.JwtService
	validator      structvalidator.Validator
	timeService    time.Time
	randomString   string.RandomString
	userRepository enuser.UserRepository
}

func NewUsecase(
	hasher crypt.Hash,
	jwtService jwt.JwtService,
	validator structvalidator.Validator,
	timeService time.Time,
	randomString string.RandomString,
	userRepository enuser.UserRepository,
) enauth.AuthUsecase {
	return &AuthUsecase{
		hasher, jwtService, validator, timeService, randomString, userRepository,
	}
}
