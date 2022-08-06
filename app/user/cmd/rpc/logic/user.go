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

func (l *UserLogic) SetUser(in *pb.SetUserRequest) (bool, error) {
	rows, err := l.model.UpdateByWhere([]string{
		"id = ?",
	}, []any{
		in.Id,
	}, []string{
		"nickname = ?",
		"avatar_url	= ?",
		"mobile = ?",
		"signature = ?",
		"status = ?",
		"updated_at = ?",
	}, []any{
		in.Nickname,
		in.AvatarUrl,
		in.Mobile,
		in.Signature,
		in.Status,
		time.Now(),
	})
	if err != nil {
		return false, err
	}
	if rows != 0 {
		return true, nil
	}
	return false, nil
}

func (l *UserLogic) Logout(id int64) (bool, error) {
	rows, err := l.model.UpdateByWhere([]string{
		"id = ?",
	}, []any{
		id,
	}, []string{
		"is_delete = ?",
	}, []any{
		1,
	})
	if err != nil {
		return false, err
	}
	if rows != 0 {
		return true, nil
	}
	return false, nil
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

func (l *UserLogic) Login(in *pb.LoginRequest) (user *model.MallUser, err error) {
	result, err := l.model.FindByWhere([]string{
		"id", "nickname", "mobile", "password", "avatar_url",
	}, []string{
		"mobile = ?",
	}, []any{
		in.Mobile,
	}, []string{})
	if err != nil {
		return nil, err
	}
	if result.ID == 0 {
		return nil, err
	}
	if result.IsDelete == 1 {
		return nil, errors.New("用户已经注销")
	}
	if !hash.PasswordVerify(in.Password, user.Password) {
		return nil, err
	}

	return &result, nil
}
