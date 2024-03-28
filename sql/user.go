package sql

import (
	"ChatDemo/model"
	"ChatDemo/sql/mysql"
	"ChatDemo/sql/redis"
)

func FindUserByID(id uint) model.User {
	user, err := redis.FindUserByID(id)
	if err == nil {
		return user
	}

	user, err = mysql.FindUserByID(id)
	if err != nil {
		return model.User{}
	}

	redis.SetUserByID(id, user)
	return user
}

func FindUserByNameAndPwd(name string, password string) model.User {
	user, err := redis.FindUserByNameAndPwd(name, password)
	if err == nil {
		return user
	}

	user, err = mysql.FindUserByNameAndPwd(name, password)
	if err != nil {
		return model.User{}
	}

	redis.SetUserByNameAndPwd(name, password, user)
	return user
}

func FindUserByName(name string) model.User {
	user, err := redis.FindUserByName(name)
	if err == nil {
		return user
	}

	user, err = mysql.FindUserByName(name)
	if err != nil {
		return model.User{}
	}

	redis.SetUserByName(name, user)
	return user
}

func FindUsersByPhone(phone string) []model.User {
	users, err := redis.FindUsersByPhone(phone)
	if err == nil {
		return users
	}

	users, err = mysql.FindUsersByPhone(phone)
	if err != nil {
		return []model.User{}
	}

	redis.SetUsersByPhone(phone, users)
	return users
}

func FindUsersByEmail(email string) []model.User {
	users, err := redis.FindUsersByEmail(email)
	if err == nil {
		return users
	}

	users, err = mysql.FindUsersByEmail(email)
	if err != nil {
		return []model.User{}
	}

	redis.SetUsersByEmail(email, users)
	return users
}

func CreateUser(user model.User) bool {
	ok := mysql.CreateUser(user)
	if !ok {
		return false
	}

	redis.DeleteUserByID(user.ID)
	redis.DeleteUserByName(user.Name)
	redis.DeleteUserByNameAndPwd(user.Name, user.PassWord)
	redis.DeleteUsersByPhone(user.Phone)
	redis.DeleteUsersByEmail(user.Email)

	return true
}

func DeleteUser(user model.User) bool {
	ok := mysql.DeleteUser(user)
	if !ok {
		return false
	}

	redis.DeleteUserByID(user.ID)
	redis.DeleteUserByName(user.Name)
	redis.DeleteUserByNameAndPwd(user.Name, user.PassWord)
	redis.DeleteUsersByPhone(user.Phone)
	redis.DeleteUsersByEmail(user.Email)

	return true
}

func UpdateUserInfo(userInfo model.UserInfo) bool {
	ok := mysql.UpdateUserInfo(userInfo)
	if !ok {
		return false
	}

	redis.DeleteUsersByPhone(userInfo.Phone)
	redis.DeleteUsersByEmail(userInfo.Email)

	return true
}

func UpdateUserPwd(user model.User) bool {
	ok := mysql.UpdateUserPwd(user)
	if !ok {
		return false
	}

	redis.DeleteUserByNameAndPwd(user.Name, user.PassWord)

	return true
}
