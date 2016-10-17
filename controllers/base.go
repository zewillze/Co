package controllers

import (
	"github.com/astaxie/beego"

)

type baseController struct {
	beego.Controller
	IsLogin bool
}



func (this *baseController) Get (){
	this.TplName = "index.html"
}

