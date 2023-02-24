// Code generated by hertz generator.

package core

import (
	"context"

	common "dousheng/router/biz/model/common"
	core "dousheng/router/biz/model/core"
	"sync/atomic"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var id int64 = 1
var name = "zhanglei"
var followCount int64 = 10
var followerCount int64 = 5
var isFollow = true

// UsersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var UsersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            &id,
		Name:          &name,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      &isFollow,
	},
}

var userIdSequence = int64(1)

// LoginHandler .
// @router /douyin/user/login/ [POST]
func LoginHandler(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserLoginResponse)

	token := *req.Username + *req.Password

	var failStatusCode int32 = 1
	var failStatusMsg = "User doesn't exist"

	if user, exist := UsersLoginInfo[token]; exist {
		resp.StatusCode = new(int32)
		resp.UserId = user.Id
		resp.Token = &token
	} else {
		resp.StatusCode = &failStatusCode
		resp.StatusMsg = &failStatusMsg
	}

	c.JSON(consts.StatusOK, resp)
}

// RegisterHandler .
// @router /douyin/user/register/ [POST]
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserRegisterResponse)

	token := *req.Username + *req.Password

	var failStatusCode int32 = 1
	var failStatusMsg = "User already exist"

	if _, exist := UsersLoginInfo[token]; exist {
		resp.StatusCode = &failStatusCode
		resp.StatusMsg = &failStatusMsg
	} else {
		newID := atomic.AddInt64(&userIdSequence, 1)
		newUser := common.User{
			Id:   &userIdSequence,
			Name: req.Username,
		}
		UsersLoginInfo[token] = newUser
		resp.StatusCode = new(int32)
		resp.UserId = &newID
		resp.Token = &token
	}

	c.JSON(consts.StatusOK, resp)
}

// UserHandler .
// @router /douyin/user/ [GET]
func UserHandler(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserResponse)

	var failStatusCode int32 = 1
	var failStatusMsg = "User doesn't exist"

	if user, exist := UsersLoginInfo[*req.Token]; exist {
		resp.StatusCode = new(int32)
		resp.User = &user
	} else {
		resp.StatusCode = &failStatusCode
		resp.StatusMsg = &failStatusMsg
	}

	c.JSON(consts.StatusOK, resp)
}
