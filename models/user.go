package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       string `orm:"pk;"`
	Password string
}

//注册
func LogUp(id, password string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Id: id}
	err := o.Read(&u)
	if err == nil {
		return 101, "用户名已存在!"
	} else {
		u.Password = password
		_, err = o.Insert(&u)
		if err == nil {
			return 100, "注册成功!"
		} else {
			return 102, "注册失败,请稍后再试!"
		}
	}
}

//登录
func Login(id, password string) (code int, msg string) {
	o := orm.NewOrm()
	u := User{Id: id}
	err := o.Read(&u)
	if err != nil {
		return 101, "用户不存在,请注册后登录!"
	} else if u.Password != password {
		return 102, "密码错误,请检查后重试!"
	}
	return 100, "登录成功!"
}
