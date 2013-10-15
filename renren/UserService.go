package renren

import (
	"encoding/json"
)

/**
 * 性别
 * @type {[type]}
 */
const (
	MALE = iota
	FEMALE
)

/**
 * 学历
 * @type {[type]}
 */
const (
	DOCTOR     = iota //枚举	博士
	COLLEGE           //枚举	本科
	GVY               //枚举	校工
	PRIMARY           //枚举	小学
	OTHER             //枚举	其他
	TEACHER           //枚举	教师
	MASTER            //枚举	硕士
	HIGHSCHOOL        //枚举	高中
	TECHNICAL         //枚举	中专技校
	JUNIOR            //枚举	初中
	SECRET            //枚举	保密
)

/**
 * 感情状态
 * @type {[type]}
 */
const (
	INLOVE        = iota //枚举	恋爱中
	SINGLE               //枚举	单身
	MARRIED              //枚举	已婚
	UNOBVIOUSLOVE        //枚举	暗恋
	DIVORCE              //枚举	离异
	ENGAGE               //枚举	订婚
	OUTLOVE              //枚举	失恋
)

type HomeTown struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

type BasicInfo struct {
	Sex      uint8  `json:"sex"`
	Birthday string `json:"birthday"`
	HomeTown `json:"homeTown"`
}

/**
 * 学校信息
 */
type School struct {
	Name                string `json:"name"`
	Year                string `json:"year"`
	EducationBackground uint8  `json:"educationBackground"`
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
type User struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Avatar       []Image `json:"avatar"`
	Star         uint8   `json:"star"`
	BasicInfo    `json:"basicInformation"`
	Education    []School `json:"education"`
	Works        []Work   `json:"work"`
	EmotionState uint8    `json:"emotionalState"`
	api          *APIRenRen
}

/**
 * get the user
 */
func (u *User) RequestUser(path string, parameters map[string]string) (ret interface{}, err error) {
	data, err := api.ApiGet(path, parameters)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal([]byte(data), &v)
	if err != nil {
		return nil, err
	}
	if m, ok := v.(map[string]interface{}); ok {
		ret = m["response"]
		return ret, nil
	}
	return nil, nil
}
