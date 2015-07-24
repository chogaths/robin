package opr_login

import (
	"errors"
	"github.com/martini-contrib/sessionauth"
)

type User struct {
	AutoID        int64
	authenticated bool
	Zone          string
}

func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

func (u *User) Login() {
	u.authenticated = true
}

func (u *User) Logout() {
	u.AutoID = 0
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.AutoID
}

// 根据id, 重新取回用户信息
func (u *User) GetById(id interface{}) error {
	if id.(int64) == 0 {
		return errors.New("no user")
	}
	return nil

}
