// Code generated by hertz generator.

package core

import (
	"context"
	"fmt"
	"path/filepath"

	handler "dousheng/router/biz/handler"
	core "dousheng/router/biz/model/core"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishActionHandler .
// @router /douyin/publish/action/ [POST]
func PublishActionHandler(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinPublishActionResponse)

	var failStatusCode int32 = 1
	var failStatusMsg = "User doesn't exist"

	if _, exist := UsersLoginInfo[*req.Token]; !exist {
		resp.StatusCode = &failStatusCode
		resp.StatusMsg = &failStatusMsg
		c.JSON(consts.StatusOK, resp)
		return
	}

	// data, err := c.FormFile("data")
	// if err != nil {
	// 	resp.StatusCode = &failStatusCode
	// 	errMsg := err.Error()
	// 	resp.StatusMsg = &errMsg
	// 	c.JSON(consts.StatusOK, resp)
	// 	return
	// }
	data := req.Data

	filename := filepath.Base(*req.Title)
	user := UsersLoginInfo[*req.Token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		resp.StatusCode = &failStatusCode
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusCode = new(int32)
	successMsg := finalName + " uploaded successfully"
	resp.StatusMsg = &successMsg

	c.JSON(consts.StatusOK, resp)
}

// PublishListHandler .
// @router /douyin/publish/list/ [GET]
func PublishListHandler(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinPublishListResponse)

	resp.StatusCode = new(int32)
	resp.VideoList = handler.DemoVideos

	c.JSON(consts.StatusOK, resp)
}
