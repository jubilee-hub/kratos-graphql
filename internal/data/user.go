package data

import (
	"context"

	"arkgo/ent"
	"arkgo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create 实现接口方法
func (r *userRepo) Create(ctx context.Context, u *biz.User) error {
	return r.data.WithTx(ctx, func(tx *ent.Tx) error {
		_, err := tx.User.Create().
			SetName(u.Name).
			Save(ctx)
		return err
	})
}

// Get 实现接口方法
func (r *userRepo) Get(ctx context.Context, id int) (*biz.User, error) {
	user, err := r.data.Ent.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *userRepo) ListUser(ctx context.Context, after *biz.Cursor, first *int, before *biz.Cursor, last *int, orderBy *biz.UserOrder, where *biz.UserWhereInput) (*biz.UserConnection, error) {
	d, err := r.data.Ent.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
	return d, err
}
