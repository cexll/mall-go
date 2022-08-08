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
	}, []any{
		id,
	}, []string{})
	if err != nil {
		return user, errors.New("query error")
	}
	return user, nil
}

func (l *UserLogic) SetUser(in *pb.SetUserRequest) (bool, error) {
	user, err := l.model.FindById(in.Id)
	if err != nil {
		return false, err
	}
	if in.Nickname != "" {
		user.Nickname = in.Nickname
	}
	if in.AvatarUrl != "" {
		user.AvatarUrl = in.AvatarUrl
	}
	if in.Status != 0 {
		user.Status = int8(in.Status)
	}
	if in.Mobile != "" {
		user.Mobile = in.Mobile
	}
	if in.Signature != "" {
		user.Signature = in.Signature
	}
	user.UpdatedAt = time.Now()
	rows, err := l.model.UpdateByWhere(&user, []string{
		"id = ?",
	}, []any{
		in.Id,
	})
	if err != nil {
		return false, nil
	}
	if rows == 0 {
		return true, nil
	}
	return true, nil
}

func (l *UserLogic) Logout(id int64) (bool, error) {
	user, err := l.model.FindById(id)
	if err != nil {
		return false, nil
	}

	user.IsDelete = 1
	user.UpdatedAt = time.Now()

	rows, err := l.model.UpdateByWhere(&user, []string{
		"id = ?",
	}, []any{
		id,
	})
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, nil
	}
	return true, nil
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

func (l *UserLogic) Login(in *pb.LoginRequest) (model.MallUser, error) {
	var user model.MallUser
	result, err := l.model.FindByWhere([]string{
		"id", "nickname", "mobile", "password", "avatar_url",
	}, []string{
		"mobile = ?",
	}, []any{
		in.Mobile,
	}, []string{})

	if err != nil {
		return user, err
	}
	if result.ID == 0 {
		return user, errors.New("用户不存在")
	}
	if result.IsDelete == 1 {
		return user, errors.New("用户已经注销")
	}
	if hash.PasswordVerify(in.Password, user.Password) {
		return user, errors.New("密码错误")
	}

	return result, nil
}
