package redis

import (
	"ChatDemo/model"
	"encoding/json"
	"fmt"
)

func GetContactList(userId uint) ([]model.Contact, error) {
	redisKey := fmt.Sprintf("UserContactInfo_ID_%d", userId)

	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return []model.Contact{}, err1
	}
	contacts := []model.Contact{}
	err2 := json.Unmarshal([]byte(res), &contacts)
	if err2 != nil {
		return []model.Contact{}, err2
	}
	return contacts, nil
}

func GetUsersIdByCommunityId(communityId uint) ([]uint, error) {
	redisKey := fmt.Sprintf("CommunityUserList_ID_%d", communityId)

	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return []uint{}, err1
	}
	usersId := []uint{}
	err2 := json.Unmarshal([]byte(res), &usersId)
	if err2 != nil {
		return []uint{}, err2
	}
	return usersId, nil
}

func SetContactList(userId uint, contacts []model.Contact) bool {
	redisKey := fmt.Sprintf("UserContactInfo_ID_%d", userId)
	contactsJson, err1 := json.Marshal(contacts)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, contactsJson, 0).Err()
	return err2 == nil
}

func SetUsersIdByCommunityId(communityId uint, usersId []uint) bool {
	redisKey := fmt.Sprintf("CommunityUserList_ID_%d", communityId)
	usersIdJson, err1 := json.Marshal(usersId)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, usersIdJson, 0).Err()
	return err2 == nil
}

func DeleteContactList(userId uint) bool {
	redisKey := fmt.Sprintf("UserContactInfo_ID_%d", userId)
	err := RDB.Del(redisKey).Err()
	return err == nil
}

func DeleteUsersIdByCommunityId(communityId uint) bool {
	redisKey := fmt.Sprintf("CommunityUserList_ID_%d", communityId)
	err := RDB.Del(redisKey).Err()
	return err == nil
}
