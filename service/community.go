package service

import (
	"ChatDemo/model"
	"ChatDemo/service/common"
	"ChatDemo/sql"

	"github.com/gin-gonic/gin"
)

func CreateCommunity(c *gin.Context) {
	community := model.Community{}
	communityName := c.Request.FormValue("name")
	communityImg := c.Request.FormValue("img")
	communityDesc := c.Request.FormValue("desc")
	userId, flag := c.Get("userId")
	if !flag {
		common.FailWithDetail(c, "没有权限", nil)
		return
	}

	community.Name = communityName
	community.OwnerId = userId.(uint)
	community.Img = communityImg
	community.Desc = communityDesc

	if !sql.CreateCommunity(community) {
		common.FailWithDetail(c, "创建群聊失败", nil)
		return
	}
	common.SuccessWithDetail(c, "创建群聊成功", nil)
}

func DeleteCommunity(c *gin.Context) {
	common.FailWithDetail(c, "删除群聊失败", nil)
}

func UpdateCommunity(c *gin.Context) {
	community := model.Community{}
	communityName := c.Request.FormValue("name")
	communityImg := c.Request.FormValue("img")
	communityDesc := c.Request.FormValue("desc")
	userId, flag := c.Get("userId")
	if !flag {
		common.FailWithDetail(c, "没有权限", nil)
		return
	}

	community0 := sql.FindCommunityByName(communityName)
	if community0.ID == 0 {
		common.FailWithDetail(c, "群聊不存在", nil)
		return
	}

	if community0.OwnerId != userId.(uint) {
		common.FailWithDetail(c, "没有权限", nil)
		return
	}

	community.Name = communityName
	community.OwnerId = userId.(uint)
	community.Img = communityImg
	community.Desc = communityDesc

	if !sql.UpdateCommunityInfo(community) {
		common.FailWithDetail(c, "更新群聊信息失败", nil)
		return
	}

	common.SuccessWithDetail(c, "更新群聊信息成功", nil)
}
