package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "test.beego.com/routers"
	_ "test.beego.com/system"
)

func main() {
	beego.Run()
}
