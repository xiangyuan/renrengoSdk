package renren

import (
	"encoding/json"
	"strings"
)

// access control
const (
	PRIVATE  = "PRIVATE"  //枚举	仅自己可见
	PUBLIC   = "PUBLIC"   //枚举	所有人可见
	PASSWORD = "PASSWORD" //	枚举	密码访问可见
	FRIEND   = "FRIEND"   //枚举	仅好友可见
)

//blog type
const (
	TYPE_API     = "TYPE_API"     //枚举	通过API发表的日志
	TYPE_RSS     = "TYPE_RSS"     //枚举	rss导入的日志
	TYPE_DEFAULT = "TYPE_DEFAULT" //枚举	web发表的日志
	TYPE_OTHER   = "TYPE_OTHER"   //枚举	其他途径发表的日志
	TYPE_WAP     = "TYPE_WAP"     //枚举	wap发表的日志
)

type Blog struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreateTime   string `json:"createTime"`
	AcessControl string `json:"accessControl"`
	ViewCount    int32  `json:"viewCount"`
	CommentCount int32  `json:"commentCount"`
	ShareCount   int32  `json:"shareCount"`
	Type         string `json:"type"`
}
type Blogs struct {
	ReBlogs []Blog `json:"response"`
}

/**
 * get the blog list
 * /v2/blog/list
 */
func (client *ApiClient) RequestBlogList(path string, parameters map[string]string) (ret interface{}, err error) {
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
	u := new(Blogs)
	if err = json.Unmarshal([]byte(data), u); err != nil {
		return nil, err
	}
	return u, nil
}

/**
 * create a blog
 * /v2/blog/create
 */
func (client *ApiClient) PutBlog(path string, parameters map[string]string) (ret interface{}, err error) {
	return nil, nil
}

/**
 * get specify blog
 * /v2/blog/get
 */
func (clien *ApiClient) GetBlog(path string, parameters map[string]string) (ret interface{}, err error) {
	return nil, nil
}
