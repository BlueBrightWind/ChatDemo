package service

import (
	"ChatDemo/model"
	"ChatDemo/service/common"
	"ChatDemo/sql"
	"fmt"
	"math/rand"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user := sql.FindUserByName(name)

	if user.ID == 0 {
		common.FailWithDetail(c, "用户不存在", nil)
		return
	}

	if !common.ValidPassword(password, user.Salt, user.PassWord) {
		common.FailWithDetail(c, "密码不正确", nil)
		return
	}

	token, err := common.ReleaseToken(user.ID, user.Name)
	if err != nil {
		common.FailWithDetail(c, "生成token失败", nil)
		return
	}

	common.SuccessWithDetail(c, "登录成功", token)
}

func IsUserNameValid(c *gin.Context) {
	name := c.Request.FormValue("name")
	user := sql.FindUserByName(name)
	if user.ID != 0 {
		common.FailWithDetail(c, "用户名已存在", nil)
		return
	}
	common.SuccessWithDetail(c, "用户名可用", nil)
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	salt := fmt.Sprintf("%06d", rand.Int31())

	user.Name = name
	user.PassWord = common.MakePassword(password, salt)
	user.Salt = salt

	if sql.CreateUser(user) {
		common.FailWithDetail(c, "新增用户失败", nil)
		return
	}
	common.SuccessWithDetail(c, "新增用户成功", nil)
}

func GetUser(c *gin.Context) {
	userId, flag := c.Get("userId")
	if !flag {
		common.FailWithDetail(c, "获取用户失败", nil)
		return
	}

	user := sql.FindUserByID(userId.(uint))
	if user.ID == 0 {
		common.FailWithDetail(c, "用户不存在", nil)
		return
	}

	userInfo := model.NewUserInfo(user)
	common.SuccessWithDetail(c, "查询成功", userInfo)
}

func DeleteUser(c *gin.Context) {
	common.FailWithDetail(c, "删除用户失败", nil)
}

func UpdateUserInfo(c *gin.Context) {
	userInfo := model.UserInfo{}
	name, _ := c.Get("userName")
	userInfo.Name = name.(string)
	userInfo.Phone = c.Request.FormValue("phone")
	userInfo.Email = c.Request.FormValue("email")
	userInfo.Avatar = c.Request.FormValue("avatar")

	_, err := govalidator.ValidateStruct(userInfo)
	if err != nil {
		common.FailWithDetail(c, "修改参数不匹配", nil)
		return
	}

	if !sql.UpdateUserInfo(userInfo) {
		common.FailWithDetail(c, "修改用户失败", nil)
		return
	}

	common.SuccessWithDetail(c, "修改用户成功", nil)
}
