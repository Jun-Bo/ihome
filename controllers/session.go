package controllers

import (
	_ "ihome/models"
	"ihome/utils"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp interface{}) {
	beego.Info("AreaController....RetData is called")
	this.Data["json"] = resp
	this.ServeJSON() //回给浏览器
}

//设置登录和注册信息
func (this *SessionController) GetName() {
	beego.Info("getName is called")
	var resp NormalResp
	resp.Errno = utils.RECODE_SESSIONERR
	resp.Errmsg = utils.RecodeText(resp.Errno)
	defer this.RetData(&resp)

	//从session里获取name字段
	name := this.GetSession("name")
	if name == nil {
		beego.Info("session name is nil")
		return
	}

	beego.Info("get name == ", name)
	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(resp.Errno)
	mapName := make(map[string]string)
	mapName["name"] = name.(string)
	resp.Data = mapName
}

func (this *SessionController) UserLogOut() {
	beego.Info("UserLogOut is called")
	var resp NormalResp
	resp.Errno = utils.RECODE_OK
	resp.Errmsg = utils.RecodeText(resp.Errno)
	defer this.RetData(&resp)

	//删除会话
	this.DelSession("user_id")
	this.DelSession("name")

}
