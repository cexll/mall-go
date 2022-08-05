package logic

import (
	"errors"
	"mall-go/app/user/cmd/pb"
	"mall-go/app/user/cmd/rpc/model"
	"mall-go/common/hash"
	"time"
)

type UserLogic struct {
	model model.MallUser
}

func (l *UserLogic) GetUser(id int64) (model.MallUser, error) {
	user, err := l.model.FindByWhere([]string{
		"id", "nickname", "avatar_url", "mobile", "signature", "status", "is_delete", "created_at", "updated_at",
	}, []string{
		"id = ?",
		"is_delete = 0",
	}, []any{id}, []string{})
	if err != nil {
		return user, errors.New("query error")
	}
	return user, nil
}

func (l *UserLogic) SetUser() {

}

func (l *UserLogic) Logout() {

}

func (l *UserLogic) Register(in *pb.RegisterRequest) (int64, error) {
	if in.Password != "" {
		hashPass, err := hash.PasswordHash(in.Password)
		if err != nil {
			return 0, err
		}
		in.Password = hashPass
	}

	return l.model.CreateOne(&model.MallUser{
		Nickname:  in.Nickname,
		Password:  in.Password,
		Mobile:    in.Mobile,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
