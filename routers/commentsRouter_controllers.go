package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddArticle",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "DeleteArticle",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetActivityArticle",
            Router: `/activity`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetCultureArticle",
            Router: `/culture`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetInformArticle",
            Router: `/inform`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetSearchArticle",
            Router: `/search/:search`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetSliderArticle",
            Router: `/slider`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ArticleController"] = append(beego.GlobalControllerRouter["travel/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetTravelArticle",
            Router: `/travel`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:CommentController"] = append(beego.GlobalControllerRouter["travel/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetComment",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:CommentController"] = append(beego.GlobalControllerRouter["travel/controllers:CommentController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:CommentController"] = append(beego.GlobalControllerRouter["travel/controllers:CommentController"],
        beego.ControllerComments{
            Method: "DeleteComment",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:ImgController"] = append(beego.GlobalControllerRouter["travel/controllers:ImgController"],
        beego.ControllerComments{
            Method: "GetSliderImg",
            Router: `/slider`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:UserController"] = append(beego.GlobalControllerRouter["travel/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["travel/controllers:UserController"] = append(beego.GlobalControllerRouter["travel/controllers:UserController"],
        beego.ControllerComments{
            Method: "LogUp",
            Router: `/logup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
