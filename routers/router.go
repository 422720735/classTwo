package routers

import (
	"classTwo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    // 注册
    beego.Router("/register", &controllers.RegController{}, "get:ShowReg;post:HandleReg")
    // 登陆
    beego.Router("/", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")


    // 文章
    beego.Router("/ShowArticle", &controllers.ArticleController{}, "get:ShowArticleList")
    // 添加文章
    beego.Router("AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleArticle")

	beego.Router("/ArticleContent", &controllers.ArticleController{}, "get:ShowArticleContent")
}
