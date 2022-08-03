package adminc_role

import (
	"context"
	"errors"
	usermodel "task1/modules/user/model"
)

type DeleteUserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
	DeleteUser(ctx context.Context, conditions map[string]interface{}) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, email string) error {

	oldUser, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return err
	}

	if oldUser.Active == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.DeleteUser(ctx, map[string]interface{}{"email": email}); err != nil {
		return err
	}

	return nil
}
