package renren

import (
	"fmt"
	"testing"
)

func TestRequestBlogList(T *testing.T) {
	param := map[string]string{
		"ownerId":    "228076041",
		"pageSize":   "10",
		"pageNumber": "1",
	}
	client := &ApiClient{
		api: api,
	}
	ret, err := client.RequestBlogList("blog/list", param)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", ret)
}

func TestPutBlog(T *testing.T) {

}

func TestGetBlog(T *testing.T) {

}
