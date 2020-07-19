package main

import (
	"fmt"
	"roblox/data_model"
	"roblox/instance"
	"roblox/classes/folder"
	"roblox/classes/model"
	// "roblox/classes/part"
	// "roblox/data_types/vector2"
)


func main() {
	game := data_model.DataModel{}
	game.Init()

	folder_1 := folder.Folder{instance.Instance{Name:"Folder_1"},}
	game.Workspace.AddChild(&folder_1)

	folder_1.AddEventListener("child_added", func(child instance.Object) {
		fmt.Println(child.GetInstance().Name)
	})

	model_1 := model.Model{instance.Instance{Name:"Model_1"},}
	folder_1.AddChild(&model_1)

	/*
	part := part.Part{Instance: instance.Instance{Name: "part_1"},}
	part.Position = vector2.Vector2{X:1, Y:6}
	part.Size = vector2.Vector2{X:10, Y:10}

	model_1.AddChild(&part)

	game.Tree()
	*/
	for {}
}
