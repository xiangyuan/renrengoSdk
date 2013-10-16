package renren

import (
	"fmt"
	"testing"
)

// var (
// 	api *APIRenRen
// )

// const (
// 	appKey      = "a5bf896d8e704e3dbeee8199d528459f"
// 	appSecret   = "877ce8975db543ac913f71f5d07cc864"
// 	redirectURL = "http://127.0.0.1:8080/oauth_callback"
// 	tmpToken    = "133922|6.3bbbe8ce73b05cb3fcfd5ae67e1e56ee.2592000.1384416000-228076041"
// )

// func init() {
// 	if api == nil {
// 		api = NewAPI(appKey, appSecret, redirectURL, "")
// 		api.SetAccessToken(tmpToken)
// 	}
// }

func TestRequestUser(t *testing.T) {
	param := map[string]string{
		"userId": "228076041",
	}
	client := &ApiClient{
		api: api,
	}
	v, err := client.RequestUser("user/get", param)
	if err != nil {
		fmt.Printf("Error %v", err.Error())
		return
	}
	fmt.Println(v)
}

func TestFriendList(t *testing.T) {
	param := map[string]string{
		"userId": "228076041",
	}
	client := &ApiClient{
		api: api,
	}
	v, err := client.RequestUser("user/friend/list", param)
	if err != nil {
		fmt.Printf("Error %v", err.Error())
		return
	}
	fmt.Println(v)
}
