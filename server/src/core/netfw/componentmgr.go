package netfw

import (
	"log"
	"sync"
)

type ComponentManager struct {

	// 组件
	componentMap      map[string]IComponent
	componentMapGuard sync.Mutex

	p *Peer

	pi IPeer
}

func (self *ComponentManager) InstallComponent(name string) IComponent {
	c := CreateComponent(name, self.pi)

	if c == nil {
		log.Printf("%s component not found: %s", self.p.GetID(), name)
		return nil
	}

	self.componentMap[name] = c

	if c == nil {
		log.Printf("%s component define not found: %s", self.p.GetID(), name)
		return nil
	}

	log.Printf("%s install component: %s", self.p.GetID(), name)
	c.Start(self.pi)

	return c
}

func (self *ComponentManager) HasComponent(name string) bool {
	if _, ok := self.componentMap[name]; ok {
		return true
	}

	return false
}

func (self *ComponentManager) GetComponent(name string) IComponent {
	if v, ok := self.componentMap[name]; ok {
		return v
	}

	return self.InstallComponent(name)
}

func NewComponentManager(p *Peer, pi IPeer) *ComponentManager {
	return &ComponentManager{
		p:            p,
		pi:           pi,
		componentMap: make(map[string]IComponent),
	}
}
