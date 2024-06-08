package enuser

import (
	"context"
)

type UserUsecase interface {
	GetDatingUserProfile(ctx context.Context, userID string) (users User, err error)
	UpgradeToPremium(ctx context.Context, userID string) (err error)
	Swipe(ctx context.Context, input SwipeUserInput) (err error)
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	FindUserByUsername(ctx context.Context, username string) (u User, err error)
	FindById(ctx context.Context, id string) (user User, err error)
	FindByIds(ctx context.Context, id []string) (user []User, err error)
	FindUsers(ctx context.Context, gender uint8, excludeUserIds []string) (user User, err error)
	Update(ctx context.Context, user *User) (err error)
	UpdateLikeCount(ctx context.Context, userID string, sType SwipeType) (err error)
}
