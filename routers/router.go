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
    beego.Router("/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleArticle")
    // 显示文章的详情
	beego.Router("/ArticleContent", &controllers.ArticleController{}, "get:ShowArticleContent")
    // 删除文章
    beego.Router("/DeleteArticle", &controllers.ArticleController{}, "get:HandleDelete")
    // 编辑文章
    beego.Router("/UpdateDetail", &controllers.ArticleController{}, "get:ShowUpdateDetail;post:HandUpdateDetail")
}
