package netfw

import (
	"log"
)

type ComponentCreator func(IPeer) IComponent

type IComponent interface {
	Start(p IPeer) // 在Peer.Start前调用
}

var componentMap map[string]ComponentCreator = make(map[string]ComponentCreator)

// 注册组件创建器
func RegisterComponent(name string, c ComponentCreator) {

	if _, ok := componentMap[name]; ok {
		log.Printf("duplicate component name: %s", name)
	}

	componentMap[name] = c
}

func CreateComponent(name string, p IPeer) IComponent {
	if v, ok := componentMap[name]; ok {
		return v(p)
	}

	return nil
}
