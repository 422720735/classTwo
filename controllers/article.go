package controllers

import (
	"classTwo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// 处理下拉框提交请求
func (this *ArticleController) HandleSelectArticle() {
	// 接受数据
	typeName := this.GetString("select")
	if typeName == "" {
		fmt.Println("不能为空")
		return
	}
	o := orm.NewOrm()
	// 查询数据
	var articles []models.Article
	// 根据id查询 过滤器查询 相当于 where                                                    // 表名__字段名            值
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles)

}

// 显示文章
func (this *ArticleController) ShowArticleList() {
	// 查询
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	var articles []models.Article
	// qs.All(&articles) // select * from Article
	// 分页查询
	pageIndex, err := strconv.Atoi(this.GetString("pageIndex"))
	if err != nil {
		pageIndex = 1
	}
	count, err := qs.Count()
	pageSize := 2 // 每页显示多少条
	start := pageSize * (pageIndex - 1)
	qs.Limit(pageSize, start).All(&articles)
	pageCount := math.Ceil(float64(count) / float64(pageSize))
	if err != nil {
		fmt.Println("查询失败")
		return
	}

	FirstPage := false
	LastPage := false
	// 首页末页显示按钮到处理
	if pageIndex == 1 {
		FirstPage = true
	}
	if pageIndex == int(pageCount) {
		LastPage = true
	}

	// 获取类型数据
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	this.Data["types"] = types
	this.Data["FirstPage"] = FirstPage
	this.Data["LastPage"] = LastPage
	// 把数据传递给视图展示
	this.Data["count"] = count
	this.Data["pageCount"] = pageCount
	this.Data["pageIndex"] = pageIndex

	this.Data["articles"] = articles
	this.TplName = "index.html"
}

// 添加文章的显示
func (this *ArticleController) ShowAddArticle() {
	// 查询文章类型视图
	o := orm.NewOrm()
	var articleType []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&articleType)
	if err != nil {
		fmt.Println("查询类型错误")
	}
	this.Data["types"] = articleType
	this.TplName = "add.html"
}

func (this *ArticleController) ShowArticleContent() {
	// 接受传递过来的id参数
	id := this.GetString("id")
	// 获取orm对象
	o := orm.NewOrm()
	// 获取查询对象

	// 传递过来的id是字符串我们需要转行成int
	// strconv包提供了简单数据类型之间的类型转换功能。可以将简单类型转换为字符串，也可以将字符串转换为其它简单类型。
	// 字符串转int：Atoi()
	// int转字符串: Itoa()
	id2, _ := strconv.Atoi(id)
	article := models.Article{Id: id2}
	// 查询
	err := o.Read(&article)
	if err != nil {
		fmt.Println("查询数据为空")
		return
	}
	// 改变我们的阅读数量
	article.Count += 1
	o.Update(&article)

	this.Data["articles"] = article
	this.TplName = "content.html"
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

	// 接受类型
	typeName:=this.GetString("select")
	if typeName == "" {
		fmt.Println("下拉框数据错误")
		return
	}
	var articeType models.ArticleType
	articeType.TypeName = typeName
	err = o.Read(&articeType, "TypeName")
	if err != nil {
		fmt.Println("获取数据类型失败 ")
		return
	}
	article.NewsArticleType = &articeType
	_, err = o.Insert(&article)
	if err != nil {
		fmt.Println("插入数据失败")
		return
	}
	// 返回视图
	this.Redirect("/ShowArticle", 302)
}

/* 删除文章 */
func (this *ArticleController) HandleDelete() {
	// 获取url传递的参数
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	article := models.Article{Id: id}
	o.Delete(&article)
	this.Redirect("/ShowArticle", 302)
}

// 编辑文章的回显
func (this *ArticleController) ShowUpdateDetail() {
	// 接受传递过来的id参数
	id := this.GetString("id")
	if id == "" {
		fmt.Println("不能为空")
		return
	}
	// 获取orm对象
	o := orm.NewOrm()
	// 获取查询对象

	// 传递过来的id是字符串我们需要转行成int
	// strconv包提供了简单数据类型之间的类型转换功能。可以将简单类型转换为字符串，也可以将字符串转换为其它简单类型。
	// 字符串转int：Atoi()
	// int转字符串: Itoa()
	id2, _ := strconv.Atoi(id)
	article := models.Article{Id: id2}
	// 查询
	err := o.Read(&article)
	if err != nil {
		fmt.Println("查询数据为空")
		return
	}
	// 改变我们的阅读数量

	//article.Count += 1
	//o.Update(&article)

	this.Data["articles"] = article
	this.TplName = "update.html"
}

// 编辑的处理数据
func (this *ArticleController) HandUpdateDetail() {
	name := this.GetString("articleName")
	content := this.GetString("content")
	id, _ := this.GetInt("id")
	if name == "" || content == "" {
		fmt.Println("更新数据失败")
		return
	}
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		fmt.Println("上传文件失败")
		return
	}
	defer f.Close()
	if h.Size > 500000 {
		fmt.Println("文件太大")
		return
	}
	// 判断文件后缀名
	ext := path.Ext(h.Filename)
	if ext != ".gif" && ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		fmt.Println("文件上传格式不正确")
		return
	}
	// 图片不重复名字
	fileName := time.Now().Format("2006-01-02 15:04:05") // 当前时间戳
	//保存
	this.SaveToFile("uploadname", "static/img/"+fileName+ext)

	// 更新文件
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		fmt.Println("要更新的文件不存在")
	}
	article.Title = name
	article.Content = content
	article.Img = "./static/img/" + fileName + ext
	_, err = o.Update(&article)
	if err != nil {
		fmt.Println("更新失败")
		return
	}
	this.Redirect("/ShowArticle", 302)
}

// 显示文章类型
func (this *ArticleController) ShowAddArticleType() {
	// 查询文章类型视图
	o := orm.NewOrm()
	var articleType []models.ArticleType
	_, err := o.QueryTable("ArticleType").All(&articleType)
	if err != nil {
		fmt.Println("查询类型错误")
	}
	this.Data["types"] = articleType
	this.TplName = "addType.html"
}

// 添加文章类型
func (this *ArticleController) HandAddArticleType() {
	// 获取数据
	typename := this.GetString("typeName")
	// 判断数据
	if typename == "" {
		fmt.Println("数据是空，不能添加")
		return
	}
	o := orm.NewOrm()
	var articleType models.ArticleType
	articleType.TypeName = typename
	_, err := o.Insert(&articleType)
	if err != nil {
		fmt.Println("插入失败")
		return
	}
	this.Redirect("/AddArticleType", 302)
}

// 删除文章
func (this *ArticleController) HandleDeleteArticleType() {
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	articleType := models.ArticleType{Id: id}
	o.Delete(&articleType)
	this.Redirect("/AddArticleType", 302)
}
