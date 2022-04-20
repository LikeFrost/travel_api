package controllers

import (
	"travel/models"

	"github.com/astaxie/beego"
)

type ImgController struct {
	beego.Controller
}

// @Title GetSliderImg
// @Description 获取轮播图
// @router /slider [get]
func (i *ImgController) GetSliderImg() {
	img, code, msg := models.GetSliderImg()
	i.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"img":  img,
	}
	i.ServeJSON()
}
