package user

import (
	"sync"
)

type UserMgr struct {
	usrMgr map[interface{}]*User
}

func (self *UserMgr) AddUser(key interface{}, u *User) {

	self.usrMgr[key] = u
}

func (self *UserMgr) RemoveUser(key interface{}) {

	delete(self.usrMgr, key)

}

func (self *UserMgr) GetUser(key interface{}) *User {

	if v, ok := self.usrMgr[key]; ok {
		return v
	}

	return nil
}

func (self *UserMgr) IterateUser(callback func(interface{}, *User) bool) {
	for k, v := range self.usrMgr {
		if !callback(k, v) {
			return
		}
	}
}

func (self *UserMgr) GetUserCount() int {

	return len(self.usrMgr)
}

func NewUserMgr() *UserMgr {
	return &UserMgr{
		usrMgr: make(map[interface{}]*User),
	}
}

type SafeUserMgr struct {
	*UserMgr

	guard sync.RWMutex
}

func NewSafeUserMgr() *SafeUserMgr {

	return &SafeUserMgr{
		UserMgr: NewUserMgr(),
	}

}

func (self *SafeUserMgr) AddUser(key interface{}, u *User) {

	self.guard.Lock()

	self.UserMgr.AddUser(key, u)

	self.guard.Unlock()

}

func (self *SafeUserMgr) RemoveUser(key interface{}) {

	self.guard.Lock()

	self.UserMgr.RemoveUser(key)

	self.guard.Unlock()

}

func (self *SafeUserMgr) GetUser(key interface{}) *User {

	self.guard.RLock()
	defer self.guard.RUnlock()

	return self.UserMgr.GetUser(key)

}

func (self *SafeUserMgr) IterateUser(callback func(interface{}, *User) bool) {
	self.guard.RLock()

	self.UserMgr.IterateUser(callback)

	self.guard.RUnlock()

}

func (self *SafeUserMgr) GetUserCount() int {

	self.guard.RLock()
	defer self.guard.RUnlock()

	return self.UserMgr.GetUserCount()

}
