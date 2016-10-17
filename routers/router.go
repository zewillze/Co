package routers

import (
	"Co/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.BlogController{})
	beego.Router("/edit_post/:Id", &controllers.BlogController{}, "get:EditPost")
	beego.Router("/manage", &controllers.BlogController{}, "get:Manage")
	beego.Router("/edit", &controllers.BlogController{}, "get:Edit;post:Save")
	beego.Router("/login", &controllers.BlogController{}, "get:LoginPage;post:Login")
}
