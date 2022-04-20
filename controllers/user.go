package controllers

import (
	"encoding/json"
	"time"
	"travel/models"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description 登录
// @router /login [post]
func (u *UserController) Login() {
	var data map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	id := data["id"].(string)
	password := data["password"].(string)
	code, msg := models.Login(id, password)
	if code == 100 {
		//创建token
		claims := make(jwt.MapClaims)
		claims["id"] = id
		claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("password"))
		if err == nil {
			u.Data["json"] = map[string]interface{}{
				"code":  code,
				"msg":   msg,
				"id":    id,
				"token": tokenString,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code":  103,
				"msg":   msg,
				"id":    id,
				"token": "生成token失败",
			}
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
		}
	}
	u.ServeJSON()
}

// @Title LogUp
// @Description 注册
// @router /logup [post]
func (u *UserController) LogUp() {
	var data map[string]interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &data)
	id := data["id"].(string)
	password := data["password"].(string)
	code, msg := models.LogUp(id, password)
	if code == 100 {
		//创建token
		claims := make(jwt.MapClaims)
		claims["id"] = id
		claims["exp"] = time.Now().Add(time.Hour * 30).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("password"))
		if err == nil {
			u.Data["json"] = map[string]interface{}{
				"code":  code,
				"msg":   msg,
				"id":    id,
				"token": tokenString,
			}
		} else {
			u.Data["json"] = map[string]interface{}{
				"code":  103,
				"msg":   msg,
				"id":    id,
				"token": "生成失败",
			}
		}
	} else {
		u.Data["json"] = map[string]interface{}{
			"code": code,
			"msg":  msg,
		}
	}
	u.ServeJSON()
}

//验证token
func (u *UserController) ParseToken() (t *jwt.Token, err string) {
	authString := u.Ctx.Input.Header("Authorization")
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
