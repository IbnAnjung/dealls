package enauth

import "context"

type AuthUsecase interface {
	RegisterUser(ctx context.Context, input RegisterInput) (output RegisterOutput, err error)
	Login(ctx context.Context, input LoginInput) (output LoginOutput, err error)
}
