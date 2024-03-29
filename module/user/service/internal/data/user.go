package data

import (
	"context"

	"mall-go/module/user/service/internal/biz"
	"mall-go/module/user/service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	Data *Data
	Log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		Data: data,
		Log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, in *biz.User) (userId int64, err error) {
	err = WithTx(ctx, r.Data.db, func(tx *model.Tx) error {
		po, err := tx.UserInfo.Create().SetNickname(in.Nickname).Save(ctx)
		if err != nil {
			return err
		}
		err = tx.UserIdentity.Create().SetUserID(po.ID).
			SetIdentityType(in.IdentityType).SetIdentifier(in.Identifier).
			SetCredential(in.Credential).Exec(ctx)
		userId = po.ID
		return err
	})
	return
}

func (r *userRepo) GetUserIdentity(ctx context.Context, identityType, identifier, credential string) (biz.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *userRepo) CreateUserIdentity(ctx context.Context, userId int64, identityType, identifier, credential string) (err error) {
	err = r.Data.db.UserIdentity.Create().SetUserID(userId).SetIdentityType(identityType).SetIdentifier(identifier).SetCredential(credential).Exec(ctx)
	return
}
