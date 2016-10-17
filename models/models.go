package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"net/url"
	"strings"
	"strconv"

)


type Post struct {
	Id 						int64
	Title 					string
 	Content                 string
	Keywords 				string `orm:"null"`
	Created 				time.Time `orm:"auto_now_add;type(datetime)"`
	Tags 					[]*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id 	 int64
	Name string
	Posts []*Post `orm:"reverse(many)"`
}

func GetAllPost() []Post{
	var posts []Post
	Posts().OrderBy("-Created").All(&posts)
	i := 0
	for _, p := range posts {
		orm.NewOrm().LoadRelated(&p, "Tags")
		posts[i] = p
		i++
	}
	return posts

}

func GetPostById(id int64) Post {

	o := orm.NewOrm()
	p := Post{Id:id}
	if created, nid, err := o.ReadOrCreate(&p, "Id"); err == nil {
		if created {
			fmt.Println("new insert an object. Id:", nid)
		}else {
			fmt.Println("get an object. Id:",nid)
			o.LoadRelated(&p, "Tags")
		}
	}
	return p
}

func Save(form url.Values) (int64, error){
	o := orm.NewOrm()

	title := strings.TrimSpace(form["post_title"][0])
	content := form["post_content"][0]
	id, e := strconv.ParseInt(form["post_id"][0], 10, 64)
	tags := form["tags"]
	if e != nil{
		id = 0
	}

	p := GetPostById(id)
	p.Content = content
	p.Title = title
	SaveTags(tags, p)


	id, err := o.Update(&p)

	if err != nil {
		beego.Alert(err)
		return -1, err
	}else {
		return id, nil
	}
}



func SaveTags(tags []string, post Post)  (int, error){

	ts := make([]*Tag, len(tags))

	i := 0
	for _,name := range tags {
		beego.Alert(name)
		tag := Tag{Name:name}
		o := orm.NewOrm()
		id, err := o.Insert(&tag)
		tag.Id = id

		if err != nil {
			beego.Alert(err)
			return -1, err
		}else {
			o.QueryM2M(&post, "Tags").Clear()
			if _, err := o.QueryM2M(&post,"Tags").Add(tag); err!=nil {
				beego.Alert(err)
			}
		}
		ts[i] = &tag
		i++
	}
	return len(ts), nil
}

func Posts() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Post))
}



func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3","./my_db.db?charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Tag))
	orm.RunSyncdb("default", false, true)
	//orm.RegisterModelWithPrefix("zz_",  new(Post))
}

