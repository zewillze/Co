package controllers
import (
	"github.com/astaxie/beego"
	"Co/models"
	"strconv"
)

type BlogController struct {
	beego.Controller
}

// Index Page
func (this *BlogController) Get () {

	//Require All Posts
	this.Data["Posts"] = models.GetAllPost()

	this.Layout = "base/base.html"
	this.TplName = "index.html"
}


func (this *BlogController) Manage() {
	this.TplName = "layout/admin.html"
}

// Add New Post Page
func (this *BlogController) Edit() {

	this.Layout = "base/base.html"
	this.TplName = "post/edit.html"
}

// Save Post
func (this *BlogController) Save() {
	this.Ctx.Request.ParseForm()

	f := this.Ctx.Request.PostForm
	_, err :=models.Save(f)

	if err != nil {
		beego.Alert(err)
		f := beego.NewFlash()
		f.Error(err.Error())
		f.Store(&this.Controller)

		this.Redirect("/", 302)
	} else {
		beego.Alert("success")
		this.Redirect("/", 302)
	}
}


// Edit Post
func (this *BlogController) EditPost() {
	id, _ := strconv.ParseInt(this.Ctx.Input.Param(":Id"), 10, 64)
	beego.Alert(models.GetPostById(id))
	this.Data["Post"] = models.GetPostById(id)
	this.Layout = "base/base.html"
	this.TplName = "post/edit.html"
}

// Login Page
func (this *BlogController) LoginPage() {
	this.Layout = "base/base.html"
	this.TplName = "login.html"
}


func (this *BlogController) Login(){
	this.Ctx.Request.ParseForm()
	f := this.Ctx.Request.PostForm
	beego.Alert(f)
	usr := f["username"][0]
	psw := f["password"][0]
	beego.Alert(usr +  psw)
	j := map[string]string{"username": usr, "status": "OK"}
	this.Data["json"] = j
	this.ServeJSON()
}