package model

type UserInfo struct {
	ID     uint
	Name   string
	Phone  string
	Email  string
	Avatar string //头像
}

func NewUserInfo(user User) UserInfo {
	return UserInfo{
		ID:     user.ID,
		Name:   user.Name,
		Phone:  user.Phone,
		Email:  user.Email,
		Avatar: user.Avatar,
	}
}
