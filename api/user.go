package api

import (
	"encoding/json"
	"github.com/qinyer/sr-client-go/context"
	"github.com/qinyer/sr-client-go/models"
	"github.com/qinyer/sr-client-go/utils"
)

type User struct {
	*context.Context
}

// NewUser 获取会员信息上报实例
func NewUser(ctx *context.Context) *User {
	return &User{ctx}
}

// Add 添加会员信息
func (user *User) Add(req *models.UserReq) (res utils.ReportCommonRes, err error) {
	response, err := utils.NewHttpClient(user.Context).HttpPostJson("/user/add_user", req)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response), &res)
	if err != nil {
		return
	}
	return
}
