package renren

import (
	"encoding/json"
	"strings"
)

const (
	SHARE  = "SHARE"  //枚举	分享
	ALBUM  = "ALBUM"  //枚举	相册
	BLOG   = "BLOG"   //枚举	日志
	STATUS = "STATUS" //枚举	状态
	PHOTO  = "PHOTO"  //枚举
)

type CommentEntity struct {
	Id           int64  `json:"id"`
	CommentType  string `json:"commentType"`
	EntryId      int64  `json:"entryId"`
	EntryOwnerId int64  `json:"entryOwnerId"`
	AuthorId     int64  `json:"authorId"`
	Content      string `json:"content"`
	PostTime     string `json:"time"`
}

type Comments struct {
	RComments []CommentEntity `json:"response"`
}

type Comment struct {
	RComment CommentEntity `json:"response"`
}

/**
 * 得到评论列表
 * https://api.renren.com/v2/comment/list
 */
func (client *ApiClient) RequestCommentList(path string, parameters map[string]string) (ret interface{}, err error) {
	data, err := client.api.ApiGet(path, parameters)
	if err != nil {
		return nil, err
	}
	e := new(APIError)
	err = json.Unmarshal([]byte(data), e)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(e.ErrorMessage.ErrorCode, "") == false {
		return nil, e
	}
	// umarshal user
	u := new(Comments)
	if err = json.Unmarshal([]byte(data), u); err != nil {
		return nil, err
	}
	return u, nil
}

/**
 * 发布评论
 * https://api.renren.com/v2/comment/put
 */
func (client *ApiClient) PutComment(path string, parameters map[string]string) (ret interface{}, err error) {
	data, err := client.api.ApiPost(path, parameters)
	if err != nil {
		return nil, err
	}
	e := new(APIError)
	err = json.Unmarshal([]byte(data), e)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(e.ErrorMessage.ErrorCode, "") == false {
		return nil, e
	}
	// umarshal user
	u := new(Comment)
	if err = json.Unmarshal([]byte(data), u); err != nil {
		return nil, err
	}
	return u, nil
}
