package user

import (
	"core/cache"
	"core/netdef"
	"svc/common/event"
)

type User struct {
	*netdef.Session
	*cache.DataSet

	Ev *event.EventDispatcher

	// GS 快捷
	//	AccountData          *model.Account
	//	CharData             *model.Char

	// GS标志
	VerifyOK    bool // 验证成功
	EnterGameOK bool // 进入游戏成功
}

//func (self *User) AddUserModel(name string) interface{} {

//	m := model.CreateModel(name)
//	if m == nil {
//		log.Printf("create model %s failed", name)
//		return nil
//	}

//	namekey := int64(util.StringHashNoCase(name))

//	self.AddData(namekey, m)

//	return m
//}

func NewUser(ses *netdef.Session) *User {
	return &User{
		Session: ses,
		DataSet: cache.NewDataSet(),
		Ev:      event.NewEventDispatcher(),
	}
}

// 进入游戏的用户
func UserHandler(ses *netdef.Session, c netdef.IPacketContext) {

	u := GetUser(ses)

	if u == nil || !u.EnterGameOK {
		c.CancelRoute()
		return
	}

	c.Map(u)
}

// 进入游戏前的用户
func AccountHandler(ses *netdef.Session, c netdef.IPacketContext) {

	u := GetUser(ses)

	if u == nil || !u.VerifyOK || u.EnterGameOK {
		c.CancelRoute()
		return
	}

	c.Map(u)
}

// 只创建了用户对象的原始用户
func RawUserHandler(ses *netdef.Session, c netdef.IPacketContext) {

	u := GetUser(ses)

	if u == nil {
		c.CancelRoute()
		return
	}

	c.Map(u)
}
