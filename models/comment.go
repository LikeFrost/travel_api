package models

import "github.com/astaxie/beego/orm"

type Comment struct {
	Id   int `orm:"pk;auto"`
	User string
	Text string `orm:"type(text)"`
	Time string
}

func GetComment() (c []*Comment, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("comment").OrderBy("-Id").All(&c)
	if err != nil {
		return c, 101, "获取评论失败"
	} else {
		return c, 100, "获取评论成功"
	}
}
func AddComment(user string, text string, time string) (code int, msg string) {
	o := orm.NewOrm()
	c := Comment{User: user, Text: text, Time: time}
	_, err := o.Insert(&c)
	if err != nil {
		return 101, "发布评论失败"
	} else {
		return 100, "发布评论成功"
	}
}
func DeleteComment(id int) (code int, msg string) {
	o := orm.NewOrm()
	c := Comment{Id: id}
	err := o.Read(&c)
	if err == nil {
		_, e := o.Delete(&c)
		if e == nil {
			return 100, "删除评论成功"
		}
	}
	return 101, "删除评论失败"
}
