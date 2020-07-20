package part

import (
	//"fmt"
	"goblox/instance"
	"goblox/data_types/vector2"
)

type Part struct {
	instance.Instance
	Position vector2.Vector2
	Size vector2.Vector2
}