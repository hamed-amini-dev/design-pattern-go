package main

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw() {
	println("Draw Circle")
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	println("Draw Rectangle")
}

func GetShape(t string) Shape {
	switch t {
	case "Circle":
		return &Circle{}
	case "Rectangle":
		return &Rectangle{}
	}
	return nil
}

func main() {
	t := "Circle"
	s := GetShape(t)
	s.Draw()
}
