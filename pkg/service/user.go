package service

import (
	"software-theory-final/pkg/modules/model"
	"software-theory-final/pkg/modules/mysql"
)

type UserService struct {
	client *mysql.Client
}

func NewUserService(client *mysql.Client) *UserService {
	return &UserService{
		client: client,
	}
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	err := u.client.DB.Create(&user).Error
	return user, err
}

func (u *UserService) GetUserById(id string) (*model.User, error) {
	user := &model.User{}
	err := u.client.DB.Where("id = ?", id).Scan(user).Error
	return user, err
}

func (u *UserService) GetUserByMailAndPwd(mail, pwd string) (*model.User, error) {
	user := &model.User{}
	err := u.client.DB.Table("users").Where("mail = ? and password = ?", mail, pwd).Scan(user).Error
	return user, err
}
