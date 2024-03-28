package mysql

import (
	"ChatDemo/model"
)

func AddContactFriendById(userId uint, friendId uint) bool {
	// 判断是否是自己
	if userId == friendId {
		return false
	}

	// 判断用户是否存在
	user := model.User{}
	DB.Where("id = ?", friendId).First(&user)
	if user.ID == 0 {
		return false
	}

	// 判断是否已经是好友
	contact := model.Contact{}
	DB.Where("owner_id =?  and target_id =? and type=1", userId, friendId).First(&contact)
	if contact.ID != 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	contact0 := model.Contact{}
	contact0.OwnerId = userId
	contact0.TargetId = friendId
	contact0.Type = 1
	if err := DB.Create(&contact0).Error; err != nil {
		tx.Rollback()
		return false
	}
	contact1 := model.Contact{}
	contact1.OwnerId = friendId
	contact1.TargetId = userId
	contact1.Type = 1
	if err := DB.Create(&contact1).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()

	return true
}

func AddContactFriendByName(userId uint, friendName string) bool {
	user := model.User{}
	DB.Where("name = ?", friendName).First(&user)
	return AddContactFriendById(userId, user.ID)
}

func DeleteContactFriendById(userId uint, friendId uint) bool {
	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	contact := model.Contact{}
	DB.Where("owner_id =?  and target_id =? and type=1", userId, friendId).First(&contact)
	if contact.ID == 0 {
		tx.Rollback()
		return false
	}
	if err := DB.Delete(&contact).Error; err != nil {
		tx.Rollback()
		return false
	}
	contact1 := model.Contact{}
	DB.Where("owner_id =?  and target_id =? and type=1", friendId, userId).First(&contact1)
	if contact1.ID == 0 {
		tx.Rollback()
		return false

	}
	if err := DB.Delete(&contact1).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DeleteContactFriendByName(userId uint, friendName string) bool {
	user := model.User{}
	DB.Where("name = ?", friendName).First(&user)
	return DeleteContactFriendById(userId, user.ID)
}

func AddContactCommunityById(userId uint, communityId uint) bool {
	// 判断群是否存在
	community := model.Community{}
	DB.Where("id = ?", communityId).First(&community)
	if community.ID == 0 {
		return false
	}

	// 判断是否已经加群
	contact := model.Contact{}
	DB.Where("owner_id =?  and target_id =? and type =2", userId, communityId).First(&contact)
	if contact.ID != 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	contact0 := model.Contact{}
	contact0.OwnerId = userId
	contact0.TargetId = communityId
	contact0.Type = 2
	if err := DB.Create(&contact0).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func AddContactCommunityByName(userId uint, communityName string) bool {
	community := model.Community{}
	DB.Where("name = ?", communityName).First(&community)
	return AddContactCommunityById(userId, community.ID)
}

func DeleteContactCommunityById(userId uint, communityId uint) bool {
	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	contact := model.Contact{}
	DB.Where("owner_id =?  and target_id =? and type=2", userId, communityId).First(&contact)
	if contact.ID == 0 {
		tx.Rollback()
		return false
	}
	if err := DB.Delete(&contact).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DeleteContactCommunityByName(userId uint, communityName string) bool {
	community := model.Community{}
	DB.Where("name = ?", communityName).First(&community)
	return DeleteContactCommunityById(userId, community.ID)
}

func GetContactList(userId uint) ([]model.Contact, error) {
	var contacts []model.Contact
	if err := DB.Where("owner_id =?", userId).Find(&contacts).Error; err != nil {
		return []model.Contact{}, err
	}
	return contacts, nil
}

func GetUsersIdByCommunityId(communityId uint) ([]uint, error) {
	var contacts []model.Contact
	var userIds []uint
	if err := DB.Where("target_id =? and type =2", communityId).Find(&contacts).Error; err != nil {
		return []uint{}, err
	}
	for _, contact := range contacts {
		userIds = append(userIds, contact.OwnerId)
	}
	return userIds, nil
}
