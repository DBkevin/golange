package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/session"
	"goblog/pkg/view"
	"net/http"
)

type AuthController struct {
}

func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {

	view.RenderSimple(w, view.D{
		"Email":    " ",
		"Password": "",
	}, "auth.login")
	session.Forget("uid")

}
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	// 2. 尝试登录
	if err := auth.Attempt(email, password); err == nil {

		// 登录成功
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 3. 失败，显示错误提示
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}
}
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{
		"User": user.User{},
	}, "auth.register")
}
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}
	//表单规则
	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		//表单不通过，重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		_user.Create()
		if _user.ID > 0 {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}

}