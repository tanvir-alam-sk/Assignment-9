package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type HomeControllar struct {
	beego.Controller
}

func (h *HomeControllar) Get() {
	h.Data["Message"] = "Hello, Beego!"
	h.TplName = "main.html"
}
