package user_service

import (
	"github.com/kainhuck/gold/models"
	"github.com/kainhuck/gold/pkg/log"
	"github.com/kainhuck/gold/pkg/util"
)

type UserRegister struct {
	Username   string `json:"username" valid:"Required;MinSize(6);MaxSize(16)"`
	Password   string `json:"password" valid:"Required;MinSize(6);MaxSize(16)"`
	RePassword string `json:"re_password" valid:"Required"`
}

func (u *UserRegister) CheckUsername() bool {
	isRegistered, err := models.IsExistUsername(u.Username)
	if err != nil {
		log.SugarLogger.Errorf("检查用户名是否存在失败, err: %v", err)
	}
	return isRegistered
}

func (u *UserRegister) CheckSamePassword() bool {
	return u.Password == u.RePassword
}

func (u *UserRegister) Save() bool {
	err := models.SaveUser(u.Username, util.EncodeMD5(u.Password))
	if err != nil {
		log.SugarLogger.Errorf("用户保存失败, err: %v", err)
		return false
	}
	return true
}
