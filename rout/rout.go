/**
 * Created with IntelliJ IDEA.
 * User: liyajie1209
 * Date: 10/11/13
 * Time: 3:53 PM
 * To change this template use File | Settings | File Templates.
 */
package rout

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sync"
	"strings"
	"regexp"
	"time"
)

const (
	CONNECT = "CONNECT"
	GET     = "GET"
	POST    = "POST"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

type Rout struct {
	name    string    `emitiy`
	pattern *regexp.Regexp
	handle  http.HandlerFunc
	method  string
	params map[int]string
}
type Router struct {
	m      sync.Mutex
	routes []*Rout // routes
	//	filters []http.HandlerFunc // options
}

func NewRouter() (r *Router) {
	return &Router{}
}

/**
 start the file server to server static resources
 */
func (r *Router) StaticResource(pattern string, resourcePath string) {
	pattern = pattern + "(.+)"

	r.AddRout(GET, pattern, func(w http.ResponseWriter, req * http.Request) {
			path := filepath.Clean(req.URL.Path)
			path = filepath.Join(resourcePath, path)
			http.ServeFile(w, req, path)
		})
}
func (r *Router) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	r.m.Lock()
	defer r.m.Unlock()
	requestPath := request.URL.Path;
	var started bool = false
    for _,route := range r.routes {
		if request.Method != route.method {
			continue
		}
		if !route.pattern.MatchString(requestPath) {
			continue
		}
		matches := route.pattern.FindStringSubmatch(requestPath)
		fmt.Println(matches)
		if len(matches[0]) != len(requestPath) {
			continue
		}
		if len(route.params) > 0 {
			values:=request.URL.Query()
			for i,value := range matches[1:] {     //将一个请求于一个特殊变量绊定
				values.Add(route.params[i],value)
			}
			fmt.Println(request.URL.RawQuery)
			if strings.EqualFold(request.URL.RawQuery,"") == false {
				request.URL.RawQuery = values.Encode() + "&" + request.URL.RawQuery
			}  else {
				request.URL.RawQuery = values.Encode()
			}
		}
		started = true
		route.handle(rw,request)
		break
		// set params
		//matches := strings.Split(match,"/")

	}
	if started == false {
		http.NotFound(rw,request)
	}
}

/*
 * add rout
 * use regex
 */
func (r *Router) AddRout(method string, pattern string, handle http.HandlerFunc) {
	r.m.Lock()
	defer r.m.Unlock()
	subPatterns := strings.Split(pattern, "/")
	params := make(map[int]string)
	j := 0
	for i, p := range subPatterns {
		// split params
		if strings.HasPrefix(p,":") {
			 expr := "([^/]+)"//匹配所有
			 //expr := ":[\\w]{1,}"   //this match :id213123 maybe have :category is params too
//			 b,err :=regexp.MatchString(expr,p)
//			 if err != nil {
//				 fmt.Println(err)
//			 }
			//如果存在id123233连接情况//后面是参数的id值
			if index:= strings.Index(p,"("); index != -1 {
				expr = p[index:]
				p = p[:index]
			}
			params[j] = p[1:]
			subPatterns[i] = expr
//		    if b,_ :=regexp.MatchString(expr,p); b == true {
//				params[i] = p[1:]
//				fmt.Println(params)
//				subPatterns[i] = expr[1:]
//
//			 }  else {
//				 params[j] = p
//				subPatterns[i]= pattern
//			 }
			j ++
		}
	}
	nPattern:= strings.Join(subPatterns,"/")
	fmt.Println(nPattern," ",params)
	regex := regexp.MustCompile(nPattern)
	route := &Rout{nPattern,regex,handle,method,params}
	r.routes = append(r.routes,route)
}

type RouteMatch struct {
	req *http.Request
	Param map[string]interface {}
	Value map[string]interface {}
	timeAlpha time.Time
}



