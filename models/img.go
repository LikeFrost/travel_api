package models

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/astaxie/beego/orm"
)

type Img struct {
	Id   int `orm:"column(id);pk;size(1);auto"`
	Path string
	Type string
}

func GetSliderImg() (img []*Img, code int, msg string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("img").Filter("Type__in", "slider").All(&img)
	for _, x := range img {
		data, e := ioutil.ReadFile(x.Path)
		if e == nil {
			x.Path = base64.StdEncoding.EncodeToString(data)
		}
	}
	if err != nil {
		return img, 101, "获取轮播图失败"
	} else {
		return img, 100, "获取轮播图成功"
	}
}
