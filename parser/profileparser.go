package parser

import (
	"fmt"
	"log"
	"regexp"

	"spider/config"
	"spider/engine"
	"spider/model"
	_ "strings"
	"time"

	"github.com/bitly/go-simplejson"
)

func ParseProfile(contents []byte) engine.ParseResult {
	var re = regexp.MustCompile(config.GetSetting("Profile::Template", 0).(string))
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		profile := parseJson(json)
		// result.Items = append(result.Items, profile)
		// 存数据库
		res, err := putToDatabase(profile)
		fmt.Println(res, err)
	}
	return result
}

//解析json数据
func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	var profile model.Profile
	obj := res.Get("objectInfo")
	profile.Nickname, _ = obj.Get("nickname").String()
	profile.Workcity, _ = obj.Get("workCityString").String()
	profile.Marriage, _ = obj.Get("marriageString").String()
	profile.Age, _ = obj.Get("age").Int()
	profile.Gender, _ = obj.Get("gender").Int()
	profile.Height, _ = obj.Get("heightString").String()
	profile.Salary, _ = obj.Get("salaryString").String()
	profile.Education, _ = obj.Get("educationString").String()
	profile.Introduce, _ = obj.Get("introduceContent").String()
	profile.Avatar, _ = obj.Get("avatarURL").String()
	profile.Last_login_time, _ = obj.Get("lastLoginTimeString").String()

	return profile
}

// 存储数据
func putToDatabase(mod model.Profile) (int64, error) {
	timestr := time.Now().Format("2006-01-02 15:04:05")
	mod.Create_time = timestr

	id, err := model.AddMovie(&mod)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, err
}
