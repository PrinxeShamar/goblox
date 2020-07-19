package vector2

import (
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func (vector *Vector2) Add(a Vector2) {
	vector.X += a.X 
	vector.Y += a.Y
}

func (vector *Vector2) Sub(a Vector2) {
	vector.X -= a.X 
	vector.Y -= a.Y
}

func (vector *Vector2) Mul(a Vector2) {
	vector.X *= a.X 
	vector.Y *= a.Y
}

func (vector *Vector2) Div(a Vector2) {
	vector.X = math.Round((vector.X / a.X) *1000)/1000
	vector.Y = math.Round((vector.X / a.X) *1000)/1000
}