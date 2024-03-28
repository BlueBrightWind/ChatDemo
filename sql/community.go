package sql

import (
	"ChatDemo/model"
	"ChatDemo/sql/mysql"
	"ChatDemo/sql/redis"
)

func CreateCommunity(community model.Community) bool {
	ok := mysql.CreateCommunity(community)
	return ok
}

func DeleteCommunity(community model.Community) bool {
	ok := mysql.DeleteCommunity(community)

	if !ok {
		return false
	}

	redis.DeleteCommunityByID(community.ID)
	redis.DeleteCommunityByName(community.Name)

	return true
}

func UpdateCommunityInfo(community model.Community) bool {
	community0 := FindCommunityByID(community.ID)
	if community0.ID == 0 {
		return false
	}

	ok := mysql.UpdateCommunityInfo(community)
	if !ok {
		return false
	}

	redis.DeleteCommunityByName(community0.Name)

	return ok
}

func FindCommunityByName(name string) model.Community {
	community, err := redis.FindCommunityByName(name)
	if err == nil {
		return community
	}

	community, err = mysql.FindCommunityByName(name)
	if err != nil {
		return model.Community{}
	}

	redis.SetCommunityByName(name, community)
	return community
}

func FindCommunityByID(id uint) model.Community {
	community, err := redis.FindCommunityByID(id)
	if err == nil {
		return community
	}

	community, err = mysql.FindCommunityByID(id)
	if err != nil {
		return model.Community{}
	}

	redis.SetCommunityByID(id, community)
	return community
}
