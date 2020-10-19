package user_service

import (
	"github.com/kainhuck/gold/models"
	"github.com/kainhuck/gold/pkg/log"
	"github.com/kainhuck/gold/pkg/util"
)

type UserLogin struct {
	Username string `json:"username" valid:"Required;MinSize(6);MaxSize(16)"`
	Password string `json:"password" valid:"Required;MinSize(6);MaxSize(16)"`
}

func (u *UserLogin) IsExistUsername() bool {
	exist, err := models.IsExistUsername(u.Username)
	if err != nil {
		log.SugarLogger.Errorf("检查用户名是否存在失败, err: %v", err)
	}
	return exist
}

func (u *UserLogin) CorrectPassword() bool {
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {
		log.SugarLogger.Errorf("通过用户名获取用户失败, err: %v", err)
		return false
	}
	return user.Password == util.EncodeMD5(u.Password)
}
