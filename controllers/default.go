package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {

	beego.Alert(c.Ctx.Request.PostForm)
	var s []string = []string {"Post", c.Ctx.Input.Header("datas")}
	f := strings.Join(s, ",")
	c.Ctx.WriteString(f)
}


