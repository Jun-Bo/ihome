package routers

import (
	"ihome/controllers"
	"strings"

	_ "ihome/models"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

//将url主页地址重定向
func ignoreStaticPath() {

	//透明static

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}

func init() {
	ignoreStaticPath() //url重定向
	beego.Router("/", &controllers.MainController{})

	//添加营业区查询路由
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetAreas")

	//设置登录注册信息
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetName")

	//处理注册功能
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:UserReg")

	//登录功能
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:UserLogin")

	//更新用户名
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "put:UpdateUserName")

	//上传头像 /api/v1.0/user/avatar
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:UploadUserPic")

}
