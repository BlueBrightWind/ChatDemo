package mysql

import (
	"ChatDemo/model"
)

func CreateCommunity(community model.Community) bool {
	// 判断群名称是否存在
	community0, _ := FindCommunityByName(community.Name)
	if community0.ID != 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Create(&community).Error; err != nil {
		return false
	}
	tx.Commit()
	return true
}

func DeleteCommunity(community model.Community) bool {
	// 判断群名称是否存在
	community0, _ := FindCommunityByName(community.Name)
	if community0.ID == 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Delete(&community).Error; err != nil {
		return false
	}
	return true
}

func UpdateCommunityInfo(community model.Community) bool {
	// 判断群名称是否存在
	community0, _ := FindCommunityByName(community.Name)
	if community0.ID == 0 {
		return false
	}

	// 开启事务
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := DB.Model(&community0).Updates(community).Error; err != nil {
		return false
	}
	tx.Commit()
	return true
}

func FindCommunityByName(name string) (model.Community, error) {
	community := model.Community{}
	if err := DB.Where("name = ?", name).First(&community).Error; err != nil {
		return model.Community{}, err
	}
	return community, nil
}

func FindCommunityByID(id uint) (model.Community, error) {
	community := model.Community{}
	if err := DB.Where("id = ?", id).First(&community).Error; err != nil {
		return model.Community{}, err
	}
	return community, nil
}
