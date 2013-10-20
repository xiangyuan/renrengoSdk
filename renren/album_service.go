package renren

import (
	"encoding/json"
	"strings"
)

//像册类型
const (
	PORTRAIT         = "PORTRAIT"         //	枚举	头像相册
	APP              = "APP"              //枚举	APP相册
	GENERAL          = "GENERAL"          //枚举	普通相册
	PUBLISHER_SINGLE = "PUBLISHER_SINGLE" //枚举	快速上传相册
	AOTHER           = "OTHER"            //枚举	其他相册
	ABLOG            = "BLOG"             //枚举	日志相册
	ALL_APP          = "ALL_APP"          //	枚举	应用相册
	VOICE            = "VOICE"            //枚举	语音相册
	HEADPHOTO        = "HEADPHOTO"        //枚举	大头贴相册
	ACTIVE           = "ACTIVE"           //枚举	活动相册
	MMS              = "MMS"              //枚举	手机相册
)

type Album struct {
	Id             int64   `json:"id"`
	Type           string  `json:"type"`
	Cover          []Image `json:"cover"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	CreateTime     string  `json:"createTime"`
	LastModifyTime string  `json:"lastModifyTime"`
	Location       string  `json:"location"`
	PhotoCount     int32   `json:"photoCount"`
	AccessControl  string  `json:"accessControl"`
}
type Albums struct {
	RAlbum []Album `json:"response"`
}

/**
 * get the blog list
 * /v2/blog/list
 */
func (client *ApiClient) RequestAlbumList(path string, parameters map[string]string) (ret interface{}, err error) {
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
	u := new(Albums)
	if err = json.Unmarshal([]byte(data), u); err != nil {
		return nil, err
	}
	return u, nil
}

/**
 * create a blog
 * /v2/blog/create
 */
func (client *ApiClient) PutAlbum(path string, parameters map[string]string) (ret interface{}, err error) {
	return nil, nil
}

/**
 * get specify blog
 * /v2/blog/get
 */
func (clien *ApiClient) GetAlbum(path string, parameters map[string]string) (ret interface{}, err error) {
	return nil, nil
}
