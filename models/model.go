package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id int
	UserName string
	Password string
}

// 文章的数据dao
type Article struct {
	Id int `orm:"pk;auto"`
	Title string`orm:"size(20)"`  // 标题
	Content string`orm:"size(500)"` // 内容
	Img string`orm:"size(50):null"` // 图片路径
	// Type string // 新闻类型
	Time time.Time`orm:"type(datetime);auto_now_add"` // 发布时间
	Count int`orm:"default(0)"` // 阅读量
}

func init()  {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/newsWeb?charset=utf8")

	orm.RegisterModel(new(User), new(Article))

	orm.RunSyncdb("default", false, true)
}
