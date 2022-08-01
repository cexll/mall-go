package logic

import (
	"errors"
	"mall-go/app/user/rpc/model"
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

func (l *UserLogic) Register() {

}
