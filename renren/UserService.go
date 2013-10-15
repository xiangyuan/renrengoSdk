package renren

import (
	"encoding/json"
)

/**
 * 性别
 * @type {[type]}
 */
const (
	MALE   = "MALE"
	FEMALE = "FEMALE"
)

/**
 * 学历
 * @type {[type]}
 */
const (
	DOCTOR     = "DOCTOR"     //枚举	博士
	COLLEGE    = "COLLEGE"    //枚举	本科
	GVY        = "GVY"        //枚举	校工
	PRIMARY    = "PRIMARY"    //枚举	小学
	OTHER      = "OTHER"      //枚举	其他
	TEACHER    = "TEACHER"    //枚举	教师
	MASTER     = "MASTER"     //枚举	硕士
	HIGHSCHOOL = "HIGHSCHOOL" //枚举	高中
	TECHNICAL  = "TECHNICAL"  //枚举	中专技校
	JUNIOR     = "JUNIOR"     //枚举	初中
	SECRET     = "SECRET"     //枚举	保密
)

/**
 * 感情状态
 * @type {[type]}
 */
const (
	INLOVE        = "INLOVE"        //枚举	恋爱中
	SINGLE        = "SINGLE"        //枚举	单身
	MARRIED       = "MARRIED"       //枚举	已婚
	UNOBVIOUSLOVE = "UNOBVIOUSLOVE" //枚举	暗恋
	DIVORCE       = "DIVORCE"       //枚举	离异
	ENGAGE        = "ENGAGE"        //枚举	订婚
	OUTLOVE       = "OUTLOVE"       //枚举	失恋
)

type HomeTown struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

type BasicInfo struct {
	Sex      string `json:"sex"`
	Birthday string `json:"birthday"`
	HomeTown `json:"homeTown"`
}

/**
 * 学校信息
 */
type School struct {
	Name                string `json:"name"`
	Year                string `json:"year"`
	EducationBackground string `json:"educationBackground"`
	Department          string `json:"department"`
}

type Industry struct {
	Category string `json:"industryCategory"`
	Detail   string `json:"industryDetail"`
	//job 忽略了
}

/**
 * 工作信息
 */
type Work struct {
	Name     string `json:"name"`
	Time     string `json:"time"`
	Industry `json:"industry"`
}

type RenRenResponse struct {
	RUser User `json:"response"`
	api   *APIRenRen
}

type User struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Avatar       []Image `json:"avatar"`
	Star         uint8   `json:"star"`
	BasicInfo    `json:"basicInformation"`
	Education    []School `json:"education"`
	Works        []Work   `json:"work"`
	EmotionState string   `json:"emotionalState"`
}

/**
 * get the user
 */
func (u *RenRenResponse) RequestUser(path string, parameters map[string]string) (ret interface{}, err error) {
	data, err := api.ApiGet(path, parameters)
	if err != nil {
		return nil, err
	}
	// var v interface{}
	err = json.Unmarshal([]byte(data), u)
	if err != nil {
		return nil, err
	}
	return u, err

	// if m, ok := v.(map[string]interface{}); ok {
	// 	ret = m["response"]
	// 	err = json.Marshal([]byte(ret))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	err = json.Unmarshal([]byte(ret), &u)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return u, nil
	// }
	// return nil, nil
}
