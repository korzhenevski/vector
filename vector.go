package vector

type Vector struct {
  X, Y, Z float64
}

func (a Vector) Add(b Vector) Vector {
  a.X += b.X
  a.Y += b.Y
  a.Z += b.Z
  return a
}

func (a Vector) Dot(b Vector) float64 {
  return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Cross(b Vector) Vector {
  return Vector{
    X: a.Y*b.Z - a.Z*b.Y,
    Y: a.Z*b.X - a.X*b.Z,
    Z: a.X*b.Y - a.Y*b.X,
  }
}

func ExampleVectorDot() {
    fmt.Println(Vector{1.0, 0.0, 0.0}.Dot(Vector{0.0, 1.0, 0.0}))
}

func ExampleVectorCross() {
    fmt.Println(Vector{1.0, 0.0, 0.0}.Cross(Vector{0.0, 1.0, 0.0}))
}
