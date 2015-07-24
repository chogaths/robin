package user

import (
	"core/netfw"
)

type IUserEvent interface {
	OnUserInit(user *User, param ...interface{}) error
}

func InvokeUserInit(componentList []string, user *User, p netfw.IPeer, param ...interface{}) error {

	// 防止重复注册
	user.Ev.Clear()

	var err error

	for _, cname := range componentList {

		// 在加载列表里没有的组件, 就不参与初始化了
		if !p.HasComponent(cname) {

			continue
		}

		ie, ok := p.GetComponent(cname).(IUserEvent)

		if !ok {
			continue
		}

		if err = ie.OnUserInit(user, param...); err != nil {
			return err
		}
	}

	return nil
}
