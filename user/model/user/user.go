package user

import (
	"github.com/jinzhu/gorm"
	"micro-go/user/model"
)

type UserInfo interface {
	Find(id uint32) (*model.User, error)
	Create(*model.User) error
	Update(*model.User, int64) (*model.User, error)
	FindByField(string, string, string) (*model.User, error)
}

type User struct {
	Dbs *gorm.DB
}

func (u *User) Find(id uint32) (*model.User, error) {
	user := &model.User{}
	user.ID = uint(id)
	if err := u.Dbs.Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User)	Create(user *model.User) error {
	if err := u.Dbs.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) 	Update(user *model.User, it int64) (*model.User, error) {
	if err := u.Dbs.Model(user).Updates(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) 	FindByField(key string, value string, fields string) (*model.User, error) {
	if len(fields) == 0 {
		fields = "*"
	}
	user := &model.User{}
	if err := u.Dbs.Select(fields).Where(key+" = ?", value).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) FindToModel(userinfo *model.User) (*model.User, error) {
	user := &model.User{}
	user = userinfo
	if err := u.Dbs.Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
