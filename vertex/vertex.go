package vertex

import (
	"fmt"
	"math"
)

type Vertex struct{ x, y, z float64 }

func NewVertex(x, y, z float64) *Vertex {
	return &Vertex{x: x, y: y, z: z}
}

func NewScaledVertex(x, y, z, s float64) *Vertex {
	return NewVertex(x*s, y*s, z*s)
}

func CopyVertex(v *Vertex) *Vertex {
	return NewVertex(v.x, v.y, v.z)
}

func (v *Vertex) Add(w *Vertex) *Vertex {
	return NewVertex(v.x+w.x, v.y+w.y, v.z+w.z)
}

func (v *Vertex) Sub(w *Vertex) *Vertex {
	return NewVertex(v.x-w.x, v.y-w.y, v.z-w.z)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vertex) Distance(w *Vertex) float64 {
	return math.Sqrt((v.x-w.x)*(v.x-w.x) + (v.y-w.y)*(v.z-w.z) + (v.z-w.z)*(v.z-w.z))
}

func (v *Vertex) Scale(s float64) {
	v.x *= s
	v.y *= s
	v.z *= s
}

func (v *Vertex) String() string {
	// fmt.Println(("String() method called")) // everytime a Print is called, this will be called
	return fmt.Sprintf(("(%.3f, %.3f, %.3f)"), v.x, v.y, v.z)
}

var origo Vertex = Vertex{0, 0, 0}

func Origo() *Vertex {
	return CopyVertex(&origo)
}
