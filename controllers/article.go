package controllers

import (
	"classTwo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// 显示文章
func (this *ArticleController) ShowArticle() {
	this.TplName = "index.html"
}

// 添加文章的显示
func (this *ArticleController) AddArticle() {
	this.TplName = "add.html"
}

// 添加文章的上传
func (this *ArticleController) HandleArticle() {
	articleName := this.GetString("articleName")
	content := this.GetString("content")
	// 获取文件上次
	f, h, err := this.GetFile("uploadname")
	defer f.Close()
	// 判断文件后缀名
	ext := path.Ext(h.Filename)
	if ext != ".gif" && ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		fmt.Println("文件上传格式不正确")
		return
	}
	// 判断大小
	if h.Size > 5000000 {
		fmt.Println("文件太大，不能上传")
		return
	}
	// 图片不重复名字
	fileName := time.Now().Format("2006-01-02 15:04:05") // 当前时间戳
	//保存
	this.SaveToFile("uploadname", "static/img/"+fileName+ext)
	if err != nil {
		fmt.Println("文件上传失败!")
	}
	fmt.Println("===========插入数据=========")
	// 插入数据
	o := orm.NewOrm()
	// 获取对象
	article := models.Article{}
	// 赋值
	article.Title = articleName
	article.Content = content
	article.Img = "static/img/" + fileName + ext
	_, err = o.Insert(&article)
	if err != nil {
		fmt.Println("插入数据失败")
		return
	}
	// 返回视图
	this.Redirect("/ShowArticle", 302)
}
