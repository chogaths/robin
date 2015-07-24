package user

import (
	"core/netdef"
	"sync"
)

var userCount int32
var userGuard sync.Mutex

const gameUserTag = "User"

// 将用户绑定在session上, 适用于客户端直连的服务器模型
func AddUser(ses *netdef.Session) *User {

	u := NewUser(ses)

	RecoverUser(u)

	return u
}

func RecoverUser(u *User) {

	userGuard.Lock()

	userCount = userCount + 1

	u.Session.SetTag(gameUserTag, u)

	userGuard.Unlock()
}

// 移除ses上的绑定
func RemoveUser(ses *netdef.Session) *User {

	tag := ses.GetTag(gameUserTag)

	if tag == nil {
		return nil
	}

	u := tag.(*User)

	userGuard.Lock()

	ses.SetTag(gameUserTag, nil)

	userCount = userCount - 1

	userGuard.Unlock()

	return u

}

func GetUser(ses *netdef.Session) *User {
	v := ses.GetTag(gameUserTag)
	if v != nil {
		return v.(*User)
	}
	return nil
}

func GetUserCount() int32 {
	return userCount
}
