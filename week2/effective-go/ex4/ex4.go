package main

import "fmt"

// TODO: Implement a function safeConvertToString(val interface{}) that:
// - Uses "comma, ok" idiom to safely check if val is a string
// - If it's a string, return the string and true
// - If not, return empty string and false
func safeConvertToString(val any) (string, bool) {
	str, ok := val.(string)
	return str, ok
}

// BONUS: Implement safeConvertToInt(val interface{}) (int, bool)
func safeConvertToInt(val any) (int, bool) {
	num, ok := val.(int)
	return num, ok
}

func main() {
	var x any = "hello"
	var y any = 42

	if str, ok := safeConvertToString(x); ok {
		fmt.Printf("Got string: %s\n", str)
	}

	if _, ok := safeConvertToString(y); !ok {
		fmt.Println("Not a string!")
	}

	if num, ok := safeConvertToInt(y); ok {
		fmt.Printf("Got int: %d\n", num)
	}

	if _, ok := safeConvertToInt(x); !ok {
		fmt.Println("Not an int!")
	}
}
