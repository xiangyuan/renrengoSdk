package renren

import (
	"fmt"
	"testing"
)

func TestRequestCommentList(T *testing.T) {
	param := map[string]string{
		"commentType":  BLOG,
		"entryOwnerId": "228076041",
		"entryId":      "493276114",
		"pageSize":     "20",
		"pageNumber":   "1",
	}
	client := &ApiClient{
		api: api,
	}
	_, err := client.RequestCommentList("comment/list", param)
	if err != nil {
		fmt.Printf("TestRequestUser %v", err.Error())
		return
	}
	// fmt.Println(v)
}

func TestPutComment(T *testing.T) {
	param := map[string]string{
		"commentType":  BLOG,
		"entryOwnerId": "228076041",
		"entryId":      "493276114",
		"content":      "小康同学",
		"targetUserId": "231245522",
	}
	client := &ApiClient{
		api: api,
	}
	v, err := client.PutComment("comment/put", param)
	if err != nil {
		fmt.Printf("TestRequestUser %v", err.Error())
		return
	}
	fmt.Println(v)
}
