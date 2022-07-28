package logic

import (
	"fmt"
	"github.com/cexll/mall-go/model"
	"github.com/cexll/mall-go/pkg/hash"
	"github.com/golang-module/carbon/v2"
)

type UserLogic struct {
	model model.UserModel
}

func (t *UserLogic) Find(Id int) *model.User {
	return t.model.FindById(Id)
}

func (t *UserLogic) CreateUser() int {
	pass, err := hash.PasswordHash("123456")
	if err != nil {
		pass = "123456"
	}
	user := model.User{
		Nickname:  "cexll",
		AvatarUrl: "1",
		Password:  pass,
		Mobile:    "18882399999",
		CreatedAt: fmt.Sprintf("%s", carbon.Now()),
		UpdatedAt: fmt.Sprintf("%s", carbon.Now()),
	}
	return int(t.model.CreateOne(&user))
}

func (t *UserLogic) CreateUsers() {

}

func (t *UserLogic) FindByWhere() (any, error) {
	user, err := t.model.FindByWhere([]string{
		"*",
	}, []string{
		"id = ?",
		"nickname = ?",
	}, []any{
		1,
		"cexll",
	}, []string{
		"order by id desc",
		"limit 1",
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
