package main

import (
	_ "example-beego/routers"
	//  "example-beego/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// controllers.CatController.GetCatImage()
	beego.Run()
}

