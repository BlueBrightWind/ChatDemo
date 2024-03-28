package redis

import (
	"ChatDemo/model"
	"encoding/json"
	"fmt"
)

func FindCommunityByName(name string) (model.Community, error) {
	redisKey := fmt.Sprintf("CommunityInfo_Name_%s", name)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return model.Community{}, err1
	}
	community := model.Community{}
	err2 := json.Unmarshal([]byte(res), &community)
	if err2 != nil {
		return model.Community{}, err2
	}
	return community, nil
}

func FindCommunityByID(id uint) (model.Community, error) {
	redisKey := fmt.Sprintf("CommunityInfo_ID_%d", id)
	res, err1 := RDB.Get(redisKey).Result()
	if err1 != nil {
		return model.Community{}, err1
	}
	community := model.Community{}
	err2 := json.Unmarshal([]byte(res), &community)
	if err2 != nil {
		return model.Community{}, err2
	}
	return community, nil
}

func SetCommunityByID(id uint, community model.Community) bool {
	redisKey := fmt.Sprintf("CommunityInfo_ID_%d", id)
	communityJson, err1 := json.Marshal(community)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, communityJson, 0).Err()
	return err2 == nil
}

func SetCommunityByName(name string, community model.Community) bool {
	redisKey := fmt.Sprintf("CommunityInfo_Name_%s", name)
	communityJson, err1 := json.Marshal(community)
	if err1 != nil {
		return false
	}
	err2 := RDB.Set(redisKey, communityJson, 0).Err()
	return err2 == nil
}

func DeleteCommunityByID(id uint) bool {
	redisKey := fmt.Sprintf("CommunityInfo_ID_%d", id)
	err := RDB.Del(redisKey).Err()
	return err == nil
}

func DeleteCommunityByName(name string) bool {
	redisKey := fmt.Sprintf("CommunityInfo_Name_%s", name)
	err := RDB.Del(redisKey).Err()
	return err == nil
}
