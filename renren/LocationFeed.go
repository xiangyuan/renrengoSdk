package renren

import (
	"encoding/json"
)

const (
	TYPE_ALL     = iota //枚举	全部类型
	TYPE_IMAGE          //枚举	照片类型
	TYPE_CHECKIN        //	枚举	签到类型
	TYPE_STATUS         //枚举	状态类型
	TYPE_POINT
)

const (
	MAIN  = iota //枚举	200pt x 600pt
	TINY         //枚举	50pt x 50pt
	LARGE        //枚举	720pt x 720pt
	HEAD         // 100
)

type Image struct {
	Size uint8  `json:"size"`
	URL  string `json:"url"`
}

type LocationPhoto struct {
	PhotoId     int64   `json:"photoId"`
	AlbumId     int64   `json:"albumId"`
	description string  `json:"description,omitempty"`
	Images      []Image `json:"images,omitempty"`
}

type LocationFeed struct {
	UserId           int64           `json:"userId"`
	UserName         string          `json:"userName"`
	HeadURL          string          `json:"headUrl"`
	placeId          string          `json:"placeId"`
	ReplyCount       int64           `json:"replyCount,omitempty"`
	UgcId            int64           `json:"ugcId"`
	Longitude        float64         `json:"longitude"`
	Latitude         float64         `json:"latitude"`
	PlaceName        string          `json:"placeName"`
	LocationFeedType uint8           `json:"locationFeedType"`
	Content          string          `json:"content,omitempty"`
	LocationPhoto    []LocationPhoto `json:"locationPhoto,omitempty"`
	api              *APIRenRen
}

/**
 *
 */
func (location *LocationFeed) RequestFeed(path string, params map[string]string) (feed interface{}, err error) {
	data, err := location.api.ApiGet(path, params)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal([]byte(data), &v)
	if err != nil {
		return nil, err
	}
	if m, ok := v.(map[string]interface{}); ok {
		feed = m["response"]
		return feed, nil
	}
	return nil, nil
}
