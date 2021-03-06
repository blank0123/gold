package user_service

import (
	"github.com/kainhuck/gold/models"
	"github.com/kainhuck/gold/pkg/log"
	"github.com/kainhuck/gold/pkg/util"
)

type UserEditPassword struct {
	Username    string `json:"username" valid:"Required"`
	Password    string `json:"password" valid:"Required"`
	NewPassword string `json:"new_password"`
}

func (u *UserEditPassword) IsExistUsername() bool {
	exist, err := models.IsExistUsername(u.Username)
	if err != nil {
		log.SugarLogger.Errorf("检查用户名是否存在失败, err: %v", err)
	}
	return exist
}
func (u *UserEditPassword) CorrectPassword() bool {
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {
		log.SugarLogger.Errorf("通过用户名获取用户失败, err: %v", err)
		return false
	}
	return user.Password == util.EncodeMD5(u.Password)
}

func (u *UserEditPassword) EditPassword() bool {
	err := models.EditPasswordByUsername(u.Username, util.EncodeMD5(u.NewPassword))
	if err != nil {
		log.SugarLogger.Errorf("修改密码失败, err: %v", err)
		return false
	}
	return true
}
