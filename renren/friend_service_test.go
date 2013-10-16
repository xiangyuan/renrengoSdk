package renren

import (
	"fmt"
	"testing"
)

func TestRequestFriendList(T *testing.T) {
	//"entryId":    "493276114"
	param := map[string]string{
		"userId":     "228076041",
		"pageSize":   "20",
		"pageNumber": "1",
	}
	client := &ApiClient{
		api: api,
	}
	v, err := client.RequestFriendList("friend/list", param)
	if err != nil {
		fmt.Printf("TestRequestUser %v", err.Error())
		return
	}
	fmt.Println(v)
}
