package martinihttp

import (
	"core/netfw"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type IMartiniAcceptor interface {
	GetInterface() *martini.ClassicMartini
}

type martiniAcceptor struct {
	*netfw.Peer

	martini *martini.ClassicMartini
}

func (self *martiniAcceptor) GetInterface() *martini.ClassicMartini {
	return self.martini
}

func (self *martiniAcceptor) GetType() string {
	return "acceptor"
}

func (self *martiniAcceptor) initSettings() {
	cfg := self.Define

	self.martini = martini.Classic()

	// 渲染器配置
	self.martini.Use(render.Renderer(render.Options{
		Directory:  cfg.GetTemplateDir(),
		Extensions: []string{".html"},
		Charset:    "UTF-8",
	}))

	// 静态文件服务
	self.martini.Use(martini.Static(cfg.GetStaticFileDir()))

}

func (self *martiniAcceptor) Start() {

	// web服务跑在另外一个线程
	go func() {
		self.martini.RunOnAddr(self.Address)
	}()
}
