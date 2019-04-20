package main

import (
	_ "classTwo/models"
	_ "classTwo/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 先加载视图函数
	beego.AddFuncMap("ShowPrePage", HandlePrePage)
	beego.AddFuncMap("ShowNextPage", HandleNextPage)
	beego.Run()
}

// 视图函数
func HandlePrePage(data int) int {
	pageIndex := data - 1
	return pageIndex
}

func HandleNextPage(data int) int {
	pageIndex := data + 1
	return pageIndex
}
