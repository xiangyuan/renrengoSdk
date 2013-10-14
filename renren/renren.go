package renren

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type APIError struct {
	ErrorCode int64
	Message   string
}

func (err *APIError) String() (str string) {
	return fmt.Sprintf("APIError ErrorMessage:%v ErrorCode %v", err.Message, err.ErrorCode)
}

type APIToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	expires      int64  `json:"expires_in"`
}

/**
 * check ApiToken is Expired
 * @return bool
 */
func (api *APIToken) isExpired() bool {
	return api.AccessToken == "" || api.RefreshToken == "" || api.expires < time.Now().Unix()
}

type APIRenRen struct {
	ApiKey         string
	ApiSecret      string
	RedirectURL    string
	ResponseType   string
	OAuth2URL      string
	AccessTokenURL string
	Version        string
}

func NewAPI(appkey string, secret string, redirect_url string, responseType string) (api *APIRenRen) {
	api = &APIRenRen{
		ApiKey:         appkey,
		ApiSecret:      secret,
		RedirectURL:    redirect_url,
		ResponseType:   responseType,
		Version:        VERSION,
		OAuth2URL:      AouthURL,
		AccessTokenURL: TokenURL,
	}
	return api
}

func (token *APIToken) SetAccessToken(atoken string, expires int64) {
	token.AccessToken = atoken
	token.expires = expires
}

/**
 * get the oathor url
 */
func (api *APIRenRen) OAuthorURL() (uri string, err error) {
	_url, err := url.Parse(api.OAuth2URL)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	v := url.Values{
		"client_id":     {api.ApiKey},
		"redirect_uri":  {api.RedirectURL},
		"response_type": {"code"},
	}
	q := v.Encode()
	fmt.Println("qcode ", q)
	if _url.RawQuery == "" {
		_url.RawQuery = q
	} else {
		_url.RawQuery = _url.RawQuery + "&" + q
	}
	fmt.Println("url =", _url.String())
	return _url.String(), nil
}

func (api *APIRenRen) GetAccessToken(code string) (result interface{}, err error) {
	_url, err := url.Parse(TokenURL)
	if err != nil {
		return nil, err
	}
	q := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {api.ApiKey},
		"client_secret": {api.ApiSecret},
		"redirect_uri":  {api.RedirectURL},
		"code":          {code},
	}
	ret, err := sendRequest(_url, GET, q)
	if err != nil {
		return nil, err
	}
	token := APIToken{}
	err = json.Unmarshal([]byte(ret), &token)
	if err != nil {
		fmt.Println(err.Error())
	}

	return token, nil
}

/**
 * send net request
 */
func sendRequest(_url *url.URL, method string, param url.Values) (result string, err error) {
	var body io.Reader
	switch method {
	case GET:
		if _url.RawQuery == "" {
			_url.RawQuery = param.Encode()
		} else {
			_url.RawQuery = _url.RawQuery + "&" + param.Encode()
		}
		body = nil
	case POST:
		body = strings.NewReader(param.Encode())
	}
	client := new(http.Client)
	request, err := http.NewRequest(method, _url.String(), body)
	if err != nil {
		return "", err
	}
	request.Header.Add("User-Agent", "renren/"+VERSION)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	switch response.Header.Get("Conent-Encoding") {
	case "gzip":
		fmt.Println("gzip")
	default:
		conent, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(conent), nil
	}
	return "", nil
}

// func (c *RenRenClient) NewRequest(method string, path string, body interface{}) (req *http.Request, err error) {
// 	rel, err := url.Parse(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u := c.BaseURL.ResolveReference(rel)
// 	buf := bytes.NewBuffer([]byte{})
// 	if body != nil {
// 		err := json.NewEncoder(buf).Encode(body)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	req, err = http.NewRequest(method, u.String(), body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Add("User-Agent", c.UserAgent)
// 	return req, nil
// }

// func (c *RenRenClient) DoRequest(req *http.Request, v interface{}) (response *http.Response, err error) {
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	response = newResponse(resp)

// 	c.Rate = response.Rate

// 	err = CheckResponse(resp)
// 	if err != nil {
// 		// even though there was an error, we still return the response
// 		// in case the caller wants to inspect it further
// 		return response, err
// 	}

// 	if v != nil {
// 		err = json.NewDecoder(resp.Body).Decode(v)
// 	}
// 	return response, err
// }
