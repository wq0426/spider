package model

import (
	"database/sql"
	"log"

	"github.com/astaxie/beego/orm"

	"spider/config"

	_ "github.com/go-sql-driver/mysql"
)

//珍爱网用户对象模型
type Profile struct {
	Nickname        string //姓名
	Workcity        string //所在城市
	Marriage        string //婚况
	Age             int    //年龄
	Gender          int    //性别
	Height          string //身高
	Salary          string //收入
	Education       string //教育
	Introduce       string //介绍
	Avatar          string //头像
	Create_time     string //创建时间
	Last_login_time string //修改时间
	Status          string //是否禁用
}

var db *sql.DB

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/spider_profile?charset=utf8")
}

//添加到数据库中
func AddMovie(profile_info *Profile) (int64, error) {
	sql := config.GetSetting("Model::AddSql", 0)
	result, err := db.Exec(sql.(string),
		profile_info.Nickname, profile_info.Workcity, profile_info.Marriage, profile_info.Age,
		profile_info.Gender, profile_info.Height, profile_info.Salary, profile_info.Education,
		profile_info.Introduce, profile_info.Avatar, profile_info.Create_time, profile_info.Last_login_time, 1)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, err
}
