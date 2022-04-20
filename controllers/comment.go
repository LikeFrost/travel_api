package controllers

import (
	"encoding/json"
	"travel/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type CommentController struct {
	beego.Controller
}

// @Title GetComment
// @Description 获取评论
// @router / [get]
func (c *CommentController) GetComment() {
	comment, code, msg := models.GetComment()
	c.Data["json"] = map[string]interface{}{
		"code":    code,
		"msg":     msg,
		"comment": comment,
	}
	c.ServeJSON()
}

// @Title AddComment
// @Description 添加评论
// @router / [post]
func (c *CommentController) AddComment() {
	var data map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	text := data["text"].(string)
	time := data["time"].(string)
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			user := claims["id"].(string)
			code, msg := models.AddComment(user, text, time)
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	c.ServeJSON()
}

// @Title DeleteComment
// @Description 添加评论
// @router /:id [delete]
func (c *CommentController) DeleteComment() {
	id, _ := c.GetInt(":id")
	token, err := c.ParseToken()
	if err != "" {
		c.Data["json"] = map[string]interface{}{
			"code": 102,
			"msg":  "token失效,请重新登录",
		}
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Data["json"] = map[string]interface{}{
				"code": 102,
				"msg":  "token失效,请重新登录",
			}
		} else {
			_ = claims["id"].(string)
			code, msg := models.DeleteComment(id)
			c.Data["json"] = map[string]interface{}{
				"code": code,
				"msg":  msg,
			}
		}
	}
	c.ServeJSON()
}

//验证token
func (c *CommentController) ParseToken() (t *jwt.Token, err string) {
	authString := c.Ctx.Input.Header("Authorization")
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
