package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"travel/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type ArticleController struct {
	beego.Controller
}

// @Title GetSliderArticle
// @Description 获取轮播图描述
// @router /slider [get]
func (a *ArticleController) GetSliderArticle() {
	article, code, msg := models.GetSliderArticle()
	a.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"article": article,
	}
	a.ServeJSON()
}

// @Title GetCultureArticle
// @Description 获取文化文章
// @router /culture [get]
func (a *ArticleController) GetCultureArticle() {
	article, code, msg := models.GetCultureArticle()
	a.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"article": article,
	}
	a.ServeJSON()
}

// @Title GetTravelArticle
// @Description 获取旅游攻略文章
// @router /travel [get]
func (a *ArticleController) GetTravelArticle() {
	article, code, msg := models.GetTravelArticle()
	a.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"article": article,
	}
	a.ServeJSON()
}

// @Title GetActivityArticle
// @Description 获取活动信息
// @router /activity [get]
func (a *ArticleController) GetActivityArticle() {
	article, code, msg := models.GetActivityArticle()
	a.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"article": article,
	}
	a.ServeJSON()
}

// @Title GetInformArticle
// @Description 获取通知信息
// @router /inform [get]
func (a *ArticleController) GetInformArticle() {
	article, code, msg := models.GetInformArticle()
	a.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"article": article,
	}
	a.ServeJSON()
}

// @Title GetInformArticle
// @Description 搜索
// @router /search/:search [get]
func (a *ArticleController) GetSearchArticle() {
	search := a.GetString(":search")
	article, code, msg := models.GetSearchArticle(search)
	a.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"article": article,
	}
	a.ServeJSON()
}

// @Title AddArticle
// @Description 添加文章
// @router / [post]
func (a *ArticleController) AddArticle() {
	articleType := a.GetString("articleType")
	title := a.GetString("title")
	text := a.GetString("text")
	location := a.GetString("location")
	time := a.GetString("time")
	token, err := a.ParseToken()
	if err != "" {
		a.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			a.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			_ = claims["id"].(string)
			md5Data := md5.New()
			md5Data.Write([]byte(articleType + title + location))
			path := "img/" + hex.EncodeToString(md5Data.Sum(nil)) + ".jpg"
			e := a.SaveToFile("img", path)
			if (e != nil) && (a.GetString("img") != "") {
				a.Data["json"] = map[string]interface{}{
					"code": 103,
					"msg":  "图片上传失败,请稍后再试",
				}
			} else {
				code, msg := models.AddArticle(articleType, title, text, path, location, time)
				a.Data["json"] = map[string]interface{}{
					"code": code,
					"msg":  msg,
				}
			}

		}
	}
	a.ServeJSON()
}

// @Title DeleteArticle
// @Description 添加文章
// @router /:id [delete]
func (a *ArticleController) DeleteArticle() {
	id, _ := a.GetInt(":id")
	token, err := a.ParseToken()
	if err != "" {
		a.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			a.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			_ = claims["id"].(string)
			code, msg := models.DeleteArticle(id)
			a.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	a.ServeJSON()
}

//验证token
func (a *ArticleController) ParseToken() (t *jwt.Token, err string) {
	authString := a.Ctx.Input.Header("Authorization")
	if authString == "" {
		return t, "token失效"
	}
	token, e := jwt.Parse(authString, func(token *jwt.Token) (interface{}, error) {
		return []byte("password"), nil
	})
	if e != nil {
		return token, "token失效"
	}
	if !token.Valid {
		return token, "token失效"
	}
	return token, ""
}
