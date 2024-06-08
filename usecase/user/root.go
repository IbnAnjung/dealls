package user

import (
	"github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/pkg/cache"
	"github.com/IbnAnjung/dealls/pkg/structvalidator"
	"github.com/IbnAnjung/dealls/pkg/time"
)

type userUsecase struct {
	validator      structvalidator.Validator
	timeService    time.Time
	cacheService   cache.CacheService
	userRepository enuser.UserRepository
}

func NewUsecase(
	validator structvalidator.Validator,
	timeService time.Time,
	cacheService cache.CacheService,
	userRepository enuser.UserRepository,
) enuser.UserUsecase {
	return &userUsecase{
		validator, timeService, cacheService, userRepository,
	}
}
