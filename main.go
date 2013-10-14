/**
 * Created with IntelliJ IDEA.
 * User: liyajie1209
 * Date: 10/13/13
 * Time: 1:57 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"oauth"
	"renren"
	"rout"
	"time"
)

const (
	appKey      = "a5bf896d8e704e3dbeee8199d528459f"
	appSecret   = "877ce8975db543ac913f71f5d07cc864"
	redirectURL = "http://127.0.0.1:8080/oauth_callback"
)

var api = renren.NewAPI(appKey, appSecret, redirectURL, "")

func rootPath(w http.ResponseWriter, req *http.Request) {
	fmt.Println("OAuthorinize")
	http.Redirect(w, req, oauthCfg.AuthCodeURL(""), http.StatusNotFound)
}

func rootPath2(w http.ResponseWriter, req *http.Request) {
	url, err := api.OAuthorURL()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.Redirect(w, req, url, http.StatusNotFound)
}

func callback(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	code = params.Get("code")
	ret, err := api.GetAccessToken(code)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("result", ret)
}

/*
 * oathonized request
 */
func request(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://api.renren.com/v2/user/get?access_token=" + aToken.AccessToken + "&userId=" + "228076041")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func main() {
	mux := rout.NewRouter()
	// mux.AddRout(rout.GET, "/gists/:gist_id/comments", rootPath)
	// mux.AddRout(rout.GET, "/oauth_callback", OAuthonizeRequest)
	mux.AddRout(rout.GET, "/oauth_callback", callback)
	// mux.AddRout(rout.GET, "/oauth", redirect)
	mux.AddRout(rout.GET, "/", rootPath2)
	mux.AddRout(rout.GET, "/user", request)
	http.Handle("/", mux)
	fmt.Println(time.Now().Unix(), "  133922")
	http.ListenAndServe(":8080", nil)
}
