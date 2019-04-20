package controllers

import (
	"classTwo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

func (this *RegController) HandleReg() {
	// 获取数据
	name := this.GetString("userName")
	passwrd := this.GetString("password")
	fmt.Println(name, passwrd)

	if name == "" || passwrd == "" {
		fmt.Println("用户或者密码不能为空")
		this.TplName = "register.html"
		return
	}
	// 插入数据

	// 1获取orm对象
	o := orm.NewOrm()
	// 2获取插入对象
	user := models.User{}
	user.UserName = name
	user.Password = passwrd
	// 3插入操作
	_, err := o.Insert(&user)
	// 4插入如果出错
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	//this.Ctx.WriteString("注册成功")
	// 使用重定向方式
	this.Redirect("/", 302)
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}

func (this *LoginController) HandleLogin() {
	name := this.GetString("userName")
	pwd := this.GetString("password")
	fmt.Println(name, pwd)
	if name == "" || pwd == "" {
		fmt.Println("不能为空")
		this.TplName = "login.html"
		return
	}
	o := orm.NewOrm()
	// 查询对象
	user := models.User{}
	// 查询
	user.UserName = name
	err := o.Read(&user, "UserName")
	fmt.Println(err, "eer=========")
	if err != nil {
		fmt.Println("查询失败")
		this.TplName = "login.html"
		return
	}
	// 判断密码是否和数据库一致
	if user.Password != pwd {
		fmt.Println("密码错误=============")
		this.TplName = "login.html"
		return
	}
	this.Redirect("/ShowArticle", 302)
}
