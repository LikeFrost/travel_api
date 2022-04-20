package models

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int `orm:"column(id);pk;auto"`
	Type     string
	Title    string
	Text     string `orm:"type(text)"`
	Img      string
	Time     string
	Location string
}

func GetSliderArticle() (a []*Article, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("Type__in", "slider").OrderBy("-Id").All(&a)
	if err != nil {
		return a, 101, "获取文字失败"
	} else {
		return a, 100, "获取文字成功"
	}
}
func GetCultureArticle() (a []*Article, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("Type__in", "culture").OrderBy("-Id").All(&a)
	for _, x := range a {
		data, e := ioutil.ReadFile(x.Img)
		if e == nil {
			x.Img = base64.StdEncoding.EncodeToString(data)
		}
	}
	if err != nil {
		return a, 101, "获取文章失败"
	} else {
		return a, 100, "获取文章成功"
	}
}
func GetTravelArticle() (a []*Article, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("Type__in", "travel").OrderBy("-Id").All(&a)
	if err != nil {
		return a, 101, "获取旅游攻略失败"
	} else {
		return a, 100, "获取旅游攻略成功"
	}
}
func GetActivityArticle() (a []*Article, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("Type__in", "activity").OrderBy("-Id").All(&a)
	for _, x := range a {
		data, e := ioutil.ReadFile(x.Img)
		if e == nil {
			x.Img = base64.StdEncoding.EncodeToString(data)
		}
	}
	if err != nil {
		return a, 101, "获取活动信息失败"
	} else {
		return a, 100, "获取旅游信息成功"
	}
}
func GetInformArticle() (a []*Article, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("article").Filter("Type__in", "inform").OrderBy("-Id").All(&a)
	if err != nil {
		return a, 101, "获取通知失败"
	} else {
		return a, 100, "获取通知成功"
	}
}
func GetSearchArticle(search string) (a []*Article, code int, msg string) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.And("Text__contains", search).Or("Title__contains", search)
	qs := o.QueryTable("article")
	qs = qs.SetCond(cond1)
	qs.All(&a)
	return a, 100, "查询成功"
}
func AddArticle(articleType, title, text, img, time, location string) (code int, msg string) {
	o := orm.NewOrm()
	a := Article{Type: articleType, Title: title, Text: text, Img: img, Time: time, Location: location}
	_, err := o.Insert(&a)
	if err != nil {
		return 101, "文章插入失败"
	} else {
		return 100, "文章插入成功"
	}
}
func DeleteArticle(id int) (code int, msg string) {
	o := orm.NewOrm()
	a := Article{Id: id}
	err := o.Read(&a)
	if err == nil {
		if a.Img != "" {
			os.Remove(a.Img)
		}
		_, e := o.Delete(&a)
		if e == nil {
			return 100, "删除文章成功"
		}
	}
	return 101, "删除文章失败"
}
