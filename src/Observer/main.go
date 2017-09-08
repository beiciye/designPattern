package main

import (
	"fmt"
)

type Observer interface {
	Name() string
	SayHello() string
}

type SubjA struct {
	obserMap map[string]Observer
}

func (s *SubjA) Register(o Observer) {
	key := o.Name()
	if s.obserMap == nil {
		s.obserMap = make(map[string]Observer, 0)
	}
	s.obserMap[key] = o
}

func (s *SubjA) RemoverRegister(o Observer) {
	key := o.Name()
	delete(s.obserMap, key)
}

func (s *SubjA) Notify() {
	for _, o := range s.obserMap {
		fmt.Println(o.SayHello())
	}
}

type ObsA struct {
	TagName string
}

func (o *ObsA) Name() string {
	return o.TagName
}

func (o *ObsA) SayHello() string {
	return "hello, i am " + o.Name()
}

func main() {
	s := &SubjA{}
	o1 := &ObsA{"jack"}
	o2 := &ObsA{"jhon"}
	o3 := &ObsA{"mike"}
	s.Register(o1)
	s.Register(o2)
	s.Register(o3)
	s.Notify()
}
