package repository

import (
	"context"
	"errors"

	"github.com/IbnAnjung/dealls/entity/enuser"
	coreerror "github.com/IbnAnjung/dealls/pkg/error"
	"github.com/IbnAnjung/dealls/pkg/orm"

	"github.com/IbnAnjung/dealls/repository/gorm/model"

	"gorm.io/gorm"
)

type userRepository struct {
	uow orm.GormUow
}

func NewGormUserRepository(
	uow orm.GormUow,
) enuser.UserRepository {
	return &userRepository{
		uow,
	}
}

func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (u enuser.User, err error) {
	m := model.MUser{}
	db := r.uow.GetDB().WithContext(ctx)

	if err = db.WithContext(ctx).Where("username = ?", username).Find(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		err = e
		return
	}

	u = m.ToEntity()

	return
}

func (r *userRepository) Create(ctx context.Context, u *enuser.User) (err error) {
	m := model.MUser{}
	m.FillFromEntity(*u)

	if err = r.uow.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			e = coreerror.NewCoreError(coreerror.CoreErrorDuplicate, "user_id already registered")
		}

		err = e
		return
	}

	u.ID = m.ID

	return nil
}

func (r *userRepository) FindById(ctx context.Context, id string) (user enuser.User, err error) {
	m := model.MUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("id = ?", id).First(&m).Error; err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "user tidak valid")
		}
		err = e
		return
	}

	user = m.ToEntity()
	return
}

func (r *userRepository) FindByIds(ctx context.Context, ids []string) (user []enuser.User, err error) {
	m := []model.MUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("id in (?)", ids).Find(&m).Error; err != nil {
		return
	}

	for _, v := range m {
		user = append(user, v.ToEntity())
	}

	return
}

func (r *userRepository) FindUsers(ctx context.Context, gender uint8, excludeUserIds []string) (user enuser.User, err error) {
	m := model.MUser{}
	if err = r.uow.GetDB().WithContext(ctx).Where("gender = ?", gender).
		Not(map[string]interface{}{"id": excludeUserIds}).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeNotFound, "no profile available")
		}
		return
	}

	user = m.ToEntity()
	return
}

func (r *userRepository) Update(ctx context.Context, user *enuser.User) (err error) {
	m := model.MUser{}
	m.FillFromEntity(*user)

	return r.uow.GetDB().WithContext(ctx).Updates(&m).Error
}

func (r *userRepository) UpdateLikeCount(ctx context.Context, userID string, sType enuser.SwipeType) (err error) {
	m := model.MUser{}
	db := r.uow.GetDB().Table(m.TableName()).Where("id = ?", userID)
	if sType == enuser.SwipeTypeLike {
		db = db.Updates(map[string]interface{}{"like_count": gorm.Expr("like_count + 1")})
	} else {
		db = db.Updates(map[string]interface{}{"like_count": gorm.Expr("like_count - 1")})
	}

	return db.Error
}
