package biz

import (
	"arkgo/ent"
	"context"
)

type User = ent.User
type UserConnection = ent.UserConnection
type UserWhereInput = ent.UserWhereInput
type CreateUserInput = ent.CreateUserInput
type UpdateUserInput = ent.UpdateUserInput
type UserOrder = ent.UserOrder

// UserRepo 在biz层定义仓储接口
type UserRepo interface {
	Create(context.Context, *User) error
	Get(context.Context, int) (*User, error)
	ListUser(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder, where *UserWhereInput) (*UserConnection, error)
	// ... 其他方法
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) ListUser(ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder, where *UserWhereInput) (*UserConnection, error) {
	return uc.repo.ListUser(ctx, after, first, before, last, orderBy, where)
}
