package controllers

import (
	//"../models"
	"runtime"
	"strings"
	"time"

	"../libs"
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

// 首页
func (this *MainController) Index() {
	this.Data["pageTitle"] = "Что-то"

	//this.Data["startJob"] = startJob
	//this.Data["okJob"] = okJob
	//this.Data["failJob"] = failJob
	//this.Data["totalJob"] = count
	//
	//this.Data["recentLogs"] = recentLogs
	//// this.Data["errLogs"] = errLogs
	//this.Data["jobs"] = jobList
	this.Data["cpuNum"] = runtime.NumCPU()
	this.display()
}

//个人信息
func (this *MainController) Profile() {
	beego.ReadFromRequest(&this.Controller)
	user := this.userService.GetByID(this.userId)

	//if this.isPost() {
	//	user.Email = this.GetString("email")
	//	user.Update()
	//	password1 := this.GetString("password1")
	//	password2 := this.GetString("password2")
	//	if password1 != "" {
	//		if len(password1) < 6 {
	//			this.ajaxMsg("密码长度必须大于6位", MSG_ERR)
	//		} else if password2 != password1 {
	//			this.ajaxMsg("两次输入的密码不一致", MSG_ERR)
	//		} else {
	//			user.Salt = string(utils.RandomCreateBytes(10))
	//			user.Password = libs.Md5([]byte(password1 + user.Salt))
	//			user.Update()
	//		}
	//	}
	//	this.ajaxMsg("", MSG_OK)
	//}

	this.Data["pageTitle"] = "资料修改"
	this.Data["user"] = user
	this.TplName = "index.tpl"
	this.display()

}

// 登录
func (this *MainController) Login() {
	if this.userId > 0 {
		this.redirect("/")
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")

		if username != "" && password != "" {
			user, err := this.userService.UserGetByName(username)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || (password != "admin") {
				errorMsg = "Ошибка автоитизации"
			} else {
				//user.LastIp = this.getClientIp()
				//user.LastLogin = time.Now().Unix()
				//models.UserUpdate(user)

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + username + password))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", string(user.ID)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", string(user.ID)+"|"+authkey, 86400)
				}
				this.redirect(beego.URLFor("MainController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("MainController.Login"))
		}
	}
	this.TplName = "public/login.html"
}

// 退出登录
func (this *MainController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("MainController.Login"))
}

// 获取系统时间
func (this *MainController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}
