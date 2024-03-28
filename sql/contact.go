package sql

import (
	"ChatDemo/model"
	"ChatDemo/sql/mysql"
	"ChatDemo/sql/redis"
)

func AddContactFriendById(userId uint, friendId uint) bool {
	ok := mysql.AddContactFriendById(userId, friendId)
	if !ok {
		return false
	}

	redis.DeleteContactList(userId)
	redis.DeleteContactList(friendId)
	return true
}

func AddContactFriendByName(userId uint, friendName string) bool {
	ok := mysql.AddContactFriendByName(userId, friendName)
	if !ok {
		return false
	}

	friendUser, err := redis.FindUserByName(friendName)
	if err == nil {
		redis.DeleteContactList(userId)
		redis.DeleteContactList(friendUser.ID)
		return true
	}

	friendUser, _ = mysql.FindUserByName(friendName)
	redis.DeleteContactList(userId)
	redis.DeleteContactList(friendUser.ID)
	return true
}

func DeleteContactFriendById(userId uint, friendId uint) bool {
	ok := mysql.DeleteContactFriendById(userId, friendId)
	if !ok {
		return false
	}

	redis.DeleteContactList(userId)
	redis.DeleteContactList(friendId)
	return true
}

func DeleteContactFriendByName(userId uint, friendName string) bool {
	ok := mysql.DeleteContactFriendByName(userId, friendName)
	if !ok {
		return false
	}

	friendUser, err := redis.FindUserByName(friendName)
	if err == nil {
		redis.DeleteContactList(userId)
		redis.DeleteContactList(friendUser.ID)
		return true
	}

	friendUser, _ = mysql.FindUserByName(friendName)
	redis.DeleteContactList(userId)
	redis.DeleteContactList(friendUser.ID)
	return true
}

func AddContactCommunityById(userId uint, communityId uint) bool {
	ok := mysql.AddContactCommunityById(userId, communityId)
	if !ok {
		return false
	}

	redis.DeleteContactList(userId)
	redis.DeleteUsersIdByCommunityId(communityId)
	return true
}

func AddContactCommunityByName(userId uint, communityName string) bool {
	ok := mysql.AddContactCommunityByName(userId, communityName)
	if !ok {
		return false
	}

	community, err := redis.FindCommunityByName(communityName)
	if err == nil {
		redis.DeleteContactList(userId)
		redis.DeleteUsersIdByCommunityId(community.ID)
		return true
	}

	community, err = mysql.FindCommunityByName(communityName)
	if err != nil {
		return false
	}

	redis.DeleteContactList(userId)
	redis.DeleteUsersIdByCommunityId(community.ID)
	return true
}

func DeleteContactCommunityById(userId uint, communityId uint) bool {
	ok := mysql.DeleteContactCommunityById(userId, communityId)
	if !ok {
		return false
	}

	redis.DeleteContactList(userId)
	redis.DeleteUsersIdByCommunityId(communityId)
	return true
}

func DeleteContactCommunityByName(userId uint, communityName string) bool {
	ok := mysql.DeleteContactCommunityByName(userId, communityName)
	if !ok {
		return false
	}

	community, err := redis.FindCommunityByName(communityName)
	if err == nil {
		redis.DeleteContactList(userId)
		redis.DeleteUsersIdByCommunityId(community.ID)
		return true
	}

	community, err = mysql.FindCommunityByName(communityName)
	if err != nil {
		return false
	}

	redis.DeleteContactList(userId)
	redis.DeleteUsersIdByCommunityId(community.ID)
	return true
}

func GetContactList(userId uint) []model.Contact {
	contacts, err := redis.GetContactList(userId)
	if err == nil {
		return contacts
	}

	contacts, err = mysql.GetContactList(userId)
	if err != nil {
		return []model.Contact{}
	}

	redis.SetContactList(userId, contacts)
	return contacts
}

func GetUsersIdByCommunityId(communityId uint) []uint {
	usersId, err := redis.GetUsersIdByCommunityId(communityId)
	if err == nil {
		return usersId
	}

	usersId, err = mysql.GetUsersIdByCommunityId(communityId)
	if err != nil {
		return []uint{}
	}

	redis.SetUsersIdByCommunityId(communityId, usersId)
	return usersId
}
