package handler

import (
	common "dousheng/router/biz/model/common"
)

var playUrl = "https://www.w3schools.com/html/movie.mp4"
var coverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
var id int64 = 1

// var isFavorite = false
var content = "Test Comment"
var createDate = "05-01"
var name = "TestUser"

// var isFollow = false

var DemoVideos = []*common.Video{
	&common.Video{
		Id:            &id,
		Author:        &DemoUser,
		PlayUrl:       &playUrl,
		CoverUrl:      &coverUrl,
		FavoriteCount: new(int64),
		CommentCount:  new(int64),
		IsFavorite:    new(bool),
	},
}

var DemoComments = []*common.Comment{
	&common.Comment{
		Id:         &id,
		User:       &DemoUser,
		Content:    &content,
		CreateDate: &createDate,
	},
}

var DemoUser = common.User{
	Id:            &id,
	Name:          &name,
	FollowCount:   new(int64),
	FollowerCount: new(int64),
	IsFollow:      new(bool),
}
