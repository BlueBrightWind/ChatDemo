package mysql

import (
	"ChatDemo/model"
)

func FindUserByID(id uint) (model.User, error) {
	user := model.User{}
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func FindUserByNameAndPwd(name string, password string) (model.User, error) {
	user := model.User{}
	if err := DB.Where("name = ? and pass_word=?", name, password).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func FindUserByName(name string) (model.User, error) {
	user := model.User{}
	if err := DB.Where("name = ?", name).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func FindUsersByPhone(phone string) ([]model.User, error) {
	users := []model.User{}
	if err := DB.Where("Phone = ?", phone).Find(&users).Error; err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func FindUsersByEmail(email string) ([]model.User, error) {
	users := []model.User{}
	if err := DB.Where("email = ?", email).First(&users).Error; err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func CreateUser(user model.User) bool {
	// 判断用户名是否存在
	user0, _ := FindUserByName(user.Name)
	if user0.ID != 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Create(&user).Error; err != nil {
		return false
	}
	tx.Commit()
	return true
}

func DeleteUser(user model.User) bool {
	// 判断用户是否存在
	user0, _ := FindUserByID(user.ID)
	if user0.ID == 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Delete(&user).Error; err != nil {
		return false
	}
	tx.Commit()
	return true
}

func UpdateUserInfo(userInfo model.UserInfo) bool {
	// 判断用户名是否存在
	user, _ := FindUserByName(userInfo.Name)
	if user.ID == 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Model(&user).Updates(model.User{Phone: userInfo.Phone, Email: userInfo.Email, Avatar: userInfo.Avatar}).Error; err != nil {
		return false
	}
	tx.Commit()
	return true
}

func UpdateUserPwd(user model.User) bool {
	// 判断用户名是否存在
	user0, _ := FindUserByID(user.ID)
	if user0.ID == 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Model(&user0).Updates(model.User{PassWord: user.PassWord}).Error; err != nil {
		return false
	}
	tx.Commit()
	return true
}
