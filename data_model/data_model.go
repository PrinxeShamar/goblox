package data_model

import (
	"fmt"
	"roblox/instance"
	"roblox/services/workspace"
	"strings"
)

type DataModel struct {
	Workspace instance.Object
	instance.Instance
}

func (data_model *DataModel) Init() {
	workspace := workspace.Workspace{
		instance.Instance{Name:"Workspace"},
	}
	data_model.Workspace = &workspace
	data_model.AddChild(&workspace)
}

func (data_model *DataModel) Tree() {
	fmt.Println("- game")
	_tree(data_model, 1)
}

func _tree(instance instance.Object, tabs int) {
	children := instance.GetInstance().Children
	if len(children) > 0 {
		for _, child := range(children) {
			fmt.Println(strings.Repeat("  ", tabs) + "- " + child.GetInstance().Name)
			_tree(child, tabs + 1)
		}
	}
}