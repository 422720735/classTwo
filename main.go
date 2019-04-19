package main

import (
	_ "classTwo/routers"
	"github.com/astaxie/beego"
	_ "classTwo/models"
)

func main() {
	beego.Run()
}

