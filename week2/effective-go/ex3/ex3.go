package main

import "fmt"

type Stringer interface {
	String() string
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return "Person: " + p.Name
}

// TODO: Implement a function processValue(value interface{}) that:
// - If it's a string, returns "String: " + value
// - If it's a Stringer, returns value.String()
// - If it's an int, returns "Number: " + string representation
// - Otherwise returns "Unknown type"

func processValue(value any) string {
	switch v := value.(type) {
	case string:
		return "String: " + v
	case int:
		return fmt.Sprintf("Number: %d", v)
	case Stringer:
		return v.String()
	default:
		return "Unknown type"
	}
}

func main() {
	// Test with different types
	fmt.Println(processValue("hello"))
	fmt.Println(processValue(Person{"Alice"}))
	fmt.Println(processValue(42))
	fmt.Println(processValue(3.14))
}
