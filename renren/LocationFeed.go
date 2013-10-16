package renren

import (
	"encoding/json"
	"strings"
)

const (
	TYPE_ALL     = "TYPE_ALL"     //枚举	全部类型
	TYPE_IMAGE   = "TYPE_IMAGE"   //枚举	照片类型
	TYPE_CHECKIN = "TYPE_CHECKIN" //	枚举	签到类型
	TYPE_STATUS  = "TYPE_STATUS"  //枚举	状态类型
	TYPE_POINT   = "TYPE_POINT"
)

const (
	MAIN  = "MAIN"  //枚举	200pt x 600pt
	TINY  = "TINY"  //枚举	50pt x 50pt
	LARGE = "LARGE" //枚举	720pt x 720pt
	HEAD  = "HEAD"  // 100
)

type Image struct {
	Size string `json:"size"`
	URL  string `json:"url"`
}

type LocationPhoto struct {
	PhotoId     int64   `json:"photoId"`
	AlbumId     int64   `json:"albumId"`
	description string  `json:"description,omitempty"`
	Images      []Image `json:"images,omitempty"`
}

type LocationFeed struct {
	feed struct {
		UserId           int64           `json:"userId"`
		UserName         string          `json:"userName"`
		HeadURL          string          `json:"headUrl"`
		placeId          string          `json:"placeId"`
		ReplyCount       int64           `json:"replyCount,omitempty"`
		UgcId            int64           `json:"ugcId"`
		Longitude        float64         `json:"longitude"`
		Latitude         float64         `json:"latitude"`
		PlaceName        string          `json:"placeName"`
		LocationFeedType string          `json:"locationFeedType"`
		Content          string          `json:"content,omitempty"`
		LocationPhoto    []LocationPhoto `json:"locationPhoto,omitempty"`
	} `json:"response"`
}

/**
 *
 */
func (location *ApiClient) RequestFeed(path string, params map[string]string) (feed interface{}, err error) {
	data, err := location.api.ApiGet(path, params)
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
	f := new(LocationFeed)
	err = json.Unmarshal([]byte(data), f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
