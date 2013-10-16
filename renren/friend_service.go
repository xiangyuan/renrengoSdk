package renren

import (
	"encoding/json"
	"strings"
)

type FriendList struct {
	UIds []int64 `json:"response"`
}

/**
 * 得到评论列表
 * https://api.renren.com/v2/friend/list
 */
func (client *ApiClient) RequestFriendList(path string, parameters map[string]string) (ret interface{}, err error) {
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
	u := new(FriendList)
	if err = json.Unmarshal([]byte(data), u); err != nil {
		return nil, err
	}
	return u, nil
}
