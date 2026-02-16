package main

import (
	"fmt"
	"math"
	"sort"
)

// TODO: Create a Shape interface with Area() float64 and Perimeter() float64
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Create Circle and Rectangle types that implement Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width * 2) + (r.Height * 2)
}

// TODO: Create Shapes type (slice of Shape) that implements:
// - sort.Interface (sort by area)
// - String() method that formats all shapes
type Shapes []Shape

func (s Shapes) Len() int {
	return len(s)
}

func (s Shapes) Less(i, j int) bool {
	return s[i].Area() < s[j].Area()
}

func (s Shapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Shapes) String() string {
	result := make([]string, len(s))
	for i, shape := range s {
		result[i] = processShape(shape)
	}
	return fmt.Sprint(result)
}

// TODO: Implement a function that accepts interface{} and uses type switch
// to handle Shape, string, and int differently
func processShape(value any) string {
	switch v := value.(type) {
	case Rectangle:
		return fmt.Sprintf("Rectangle(%.2fx%.2f, area=%.2f)", v.Height, v.Width, v.Area())
	case Circle:
		return fmt.Sprintf("Circle(r=%.2f, area=%.2f)", v.Radius, v.Area())
	default:
		return "Unknown shape"
	}
}

func main() {
	shapes := Shapes{
		Circle{5},
		Rectangle{4, 6},
		Circle{3},
		Rectangle{2, 8},
	}

	// Sort and display
	sort.Sort(shapes)
	fmt.Println(shapes)

	// Test type switching
	processShape(shapes[0])
	processShape("hello")
	processShape(42)
}
