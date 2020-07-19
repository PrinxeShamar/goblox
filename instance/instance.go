package instance

import (
	//"fmt"
)

type Object interface {
	GetInstance() *Instance
	AddChild(child Object) 
	RemoveChild(child Object)
	AddEventListener(event string, callback function)
}

type Instance struct {
	Name string
	Parent Object
	Children []Object
	events map[string][]function
}

func (instance *Instance) AddChild(child Object) {
	child_instance := child.GetInstance()
	if child_instance != instance && child_instance.Parent != instance {
		if child_instance.Parent != nil {
			child_instance.Parent.RemoveChild(child)
		}
		child_instance.Parent = instance
		instance.Children = append(instance.Children, child)
		instance.emit("child_added", child)
	}
}

func (instance *Instance) RemoveChild(child Object) {
	for index, _child := range(instance.Children) {
		if _child == child {
			instance.Children = append(instance.Children[:index], instance.Children[index+1:]...)
			return 
		}
	}
}

func (instance *Instance) GetInstance() *Instance {
	return instance
}

type function func(Object)

func (instance *Instance) AddEventListener(event string, callback function) {
	
	instance.events = make(map[string][]function)
	if instance.events[event] == nil {
		var callbacks []function
		instance.events[event] = callbacks
	}
	instance.events[event] = append(instance.events[event], callback)
}

func (instance *Instance) emit(event string, object Object) {
	if instance.events[event] != nil {
		for _, fn := range(instance.events[event]) {
			go fn(object)
		}
	}
}