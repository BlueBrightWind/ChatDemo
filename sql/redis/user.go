package redis

import (
	"ChatDemo/model"
	"encoding/json"
	"fmt"
)

func FindUserByID(id uint) (model.User, error) {
	redisKey := fmt.Sprintf("UserInfo_ID_%d", id)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return model.User{}, err1
	}
	user := model.User{}
	err2 := json.Unmarshal([]byte(res), &user)
	if err2 != nil {
		return model.User{}, err2
	}
	return user, nil
}

func FindUserByNameAndPwd(name string, password string) (model.User, error) {
	redisKey := fmt.Sprintf("UserInfo_ID_%s_Password_%s", name, password)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return model.User{}, err1
	}
	user := model.User{}
	err2 := json.Unmarshal([]byte(res), &user)
	if err2 != nil {
		return model.User{}, err2
	}
	return user, nil
}

func FindUserByName(name string) (model.User, error) {
	redisKey := fmt.Sprintf("UserInfo_Name_%s", name)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return model.User{}, err1
	}
	user := model.User{}
	err2 := json.Unmarshal([]byte(res), &user)
	if err2 != nil {
		return model.User{}, err2
	}
	return user, nil
}

func FindUsersByPhone(phone string) ([]model.User, error) {
	redisKey := fmt.Sprintf("UserInfo_Phone_%s", phone)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return []model.User{}, err1
	}
	users := []model.User{}
	err2 := json.Unmarshal([]byte(res), &users)
	if err2 != nil {
		return []model.User{}, err2
	}
	return users, nil
}

func FindUsersByEmail(email string) ([]model.User, error) {
	redisKey := fmt.Sprintf("UserInfo_Email_%s", email)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return []model.User{}, err1
	}
	users := []model.User{}
	err2 := json.Unmarshal([]byte(res), &users)
	if err2 != nil {
		return []model.User{}, err2
	}
	return users, nil
}

func SetUserByID(id uint, user model.User) bool {
	redisKey := fmt.Sprintf("UserInfo_ID_%d", id)
	userInfoJson, err1 := json.Marshal(user)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, userInfoJson, 0)
	return err2 == nil
}

func SetUserByNameAndPwd(name string, password string, user model.User) bool {
	redisKey := fmt.Sprintf("UserInfo_Name_%s_Password_%s", name, password)
	userInfoJson, err1 := json.Marshal(user)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, userInfoJson, 0)
	return err2 == nil
}

func SetUserByName(name string, user model.User) bool {
	redisKey := fmt.Sprintf("UserInfo_Name_%s", name)
	userInfoJson, err1 := json.Marshal(user)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, userInfoJson, 0)
	return err2 == nil
}

func SetUsersByPhone(phone string, users []model.User) bool {
	redisKey := fmt.Sprintf("UserInfo_Phone_%s", phone)
	usersJson, err1 := json.Marshal(users)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, usersJson, 0)
	return err2 == nil
}

func SetUsersByEmail(email string, users []model.User) bool {
	redisKey := fmt.Sprintf("UserInfo_Email_%s", email)
	usersJson, err1 := json.Marshal(users)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, usersJson, 0)
	return err2 == nil
}

func DeleteUserByID(id uint) bool {
	redisKey := fmt.Sprintf("UserInfo_ID_%d", id)
	err := RDB.Del(redisKey)
	return err == nil
}

func DeleteUserByNameAndPwd(name string, password string) bool {
	redisKey := fmt.Sprintf("UserInfo_Name_%s_Password_%s", name, password)
	err := RDB.Del(redisKey)
	return err == nil
}

func DeleteUserByName(name string) bool {
	redisKey := fmt.Sprintf("UserInfo_Name_%s", name)
	err := RDB.Del(redisKey)
	return err == nil
}

func DeleteUsersByPhone(phone string) bool {
	redisKey := fmt.Sprintf("UserInfo_Phone_%s", phone)
	err := RDB.Del(redisKey)
	return err == nil
}

func DeleteUsersByEmail(email string) bool {
	redisKey := fmt.Sprintf("UserInfo_Email_%s", email)
	err := RDB.Del(redisKey)
	return err == nil
}
