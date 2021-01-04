package api

import (
	"encoding/json"
	"sr-client-go/context"
	"sr-client-go/models"
	"sr-client-go/utils"
)

type User struct {
	*context.Context
}

// NewUser 获取会员信息上报实例
func NewUser(ctx *context.Context) *User {
	user := new(User)
	user.Context = ctx
	return user
}

// Add 添加会员信息
func (user *User) Add(req *models.UserReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(user.AppId, user.AppSecret, user.Prod).HttpPostJson("/user/add_user", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
