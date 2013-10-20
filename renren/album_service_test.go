package renren

import (
	"fmt"
	"testing"
)

func TestRequestAlbumList(T *testing.T) {
	param := map[string]string{
		"ownerId":    "228076041",
		"pageSize":   "10",
		"pageNumber": "1",
	}
	client := &ApiClient{
		api: api,
	}
	ret, err := client.RequestAlbumList("album/list", param)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
}

func TestPutAlbum(T *testing.T) {

}

func TestGetAlbum(T *testing.T) {

}
