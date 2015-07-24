package opr_login

import (
	"core/martinihttp"
	"core/netfw"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"net/http"
	"protos/coredef"
)

// 登陆页面post上来的数据
type gmUserPostForm struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type oprloginComponent struct {
}

func (self *oprloginComponent) Start(peer netfw.IPeer) {

	m := peer.(martinihttp.IMartiniAcceptor).GetInterface()

	var config coredef.OperateConfig
	netfw.GetConfig("OperateConfig", &config)

	store := sessions.NewCookieStore([]byte("secret123"))

	store.Options(sessions.Options{
		MaxAge: 0,
	})

	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))

	// 显示登陆页面
	m.Get("/login", func(r render.Render) {
		r.HTML(200, "login", nil)
	})

	// 提交登陆
	m.Post("/login", binding.Bind(gmUserPostForm{}), func(session sessions.Session, msg gmUserPostForm, r render.Render, req *http.Request) {

		var verify bool

		for _, v := range config.GetAccount() {
			if msg.UserName == v.GetAccount() && msg.Password == EncodePassword(v.GetPassword()) {
				verify = true
				break
			}
		}

		if !verify {
			r.Redirect("login")
		}

		err := sessionauth.AuthenticateSession(session, &User{AutoID: 1})
		if err != nil {
			r.JSON(500, err)
		}

		r.Redirect("index")

	})

	// 登出
	m.Get("/logout", sessionauth.LoginRequired, func(session sessions.Session, user sessionauth.User, r render.Render) {
		sessionauth.Logout(session, user)
		r.Redirect("/login")
	})

	m.Get("/index", sessionauth.LoginRequired, func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Get("/", func(r render.Render) {
		r.Redirect("index")
	})

}

func init() {

	netfw.RegisterComponent("login", func(p netfw.IPeer) netfw.IComponent {
		return &oprloginComponent{}
	})
}
