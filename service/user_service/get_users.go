package user_service

import "github.com/kainhuck/gold/models"

type UserGets struct {
	PageNum  int
	PageSize int
}

func (u *UserGets)Get() ([]models.User, error){
	return models.GetUsers(u.PageNum, u.PageSize)
}
