package service

import (
	"ChatDemo/model"
	"ChatDemo/service/common"
	"ChatDemo/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddContactFriend(c *gin.Context) {
	userId, _ := c.Get("userId")
	friendName := c.Request.FormValue("name")
	if !sql.AddContactFriendByName(userId.(uint), friendName) {
		common.FailWithDetail(c, "添加好友失败", nil)
		return
	}
	common.SuccessWithDetail(c, "添加好友成功", nil)
}

func RemoveContactFriend(c *gin.Context) {
	userId, _ := c.Get("userId")
	friendName := c.Request.FormValue("name")
	if !sql.DeleteContactFriendByName(userId.(uint), friendName) {
		common.FailWithDetail(c, "删除好友失败", nil)
		return
	}
	common.SuccessWithDetail(c, "删除好友成功", nil)
}

func SearchFriend(c *gin.Context) {
	name := c.Request.FormValue("name")
	user := sql.FindUserByName(name)
	if user.ID == 0 {
		common.FailWithDetail(c, "用户不存在", nil)
		return
	}

	userInfo := model.NewUserInfo(user)
	common.SuccessWithDetail(c, "查询成功", userInfo)
}

func AddContactCommunity(c *gin.Context) {
	userId, _ := c.Get("userId")
	communityName := c.Request.FormValue("name")
	if !sql.AddContactCommunityByName(userId.(uint), communityName) {
		common.FailWithDetail(c, "添加群聊失败", nil)
		return
	}
	common.SuccessWithDetail(c, "添加群聊成功", nil)
}

func RemoveContactCommunity(c *gin.Context) {
	userId, _ := c.Get("userId")
	communityName := c.Request.FormValue("name")
	if !sql.DeleteContactCommunityByName(userId.(uint), communityName) {
		common.FailWithDetail(c, "删除群聊失败", nil)
		return
	}
	common.SuccessWithDetail(c, "删除群聊成功", nil)
}

func SearchCommunity(c *gin.Context) {
	name := c.Request.FormValue("name")
	community := sql.FindCommunityByName(name)
	if community.ID == 0 {
		common.FailWithDetail(c, "群聊不存在", nil)
		return
	}

	communityInfo := model.NewCommunityInfo(community)
	common.SuccessWithDetail(c, "查询成功", communityInfo)
}

func GetContactList(c *gin.Context) {
	userId, _ := c.Get("userId")
	list := sql.GetContactList(userId.(uint))
	infoList := model.NewContactInfoList(list)
	common.SuccessWithDetail(c, "查询成功", infoList)
}

func GetCommunityInfoById(c *gin.Context) {
	idstr := c.Request.FormValue("id")
	id, _ := strconv.Atoi(idstr)
	community := sql.FindCommunityByID(uint(id))
	if community.ID == 0 {
		common.FailWithDetail(c, "群聊不存在", nil)
		return
	}
	communityInfo := model.NewCommunityInfo(community)
	common.SuccessWithDetail(c, "查询成功", communityInfo)
}

func GetUserInfoById(c *gin.Context) {
	idstr := c.Request.FormValue("id")
	id, _ := strconv.Atoi(idstr)
	user := sql.FindUserByID(uint(id))
	if user.ID == 0 {
		common.FailWithDetail(c, "用户不存在", nil)
		return
	}
	userInfo := model.NewUserInfo(user)
	common.SuccessWithDetail(c, "查询成功", userInfo)
}
