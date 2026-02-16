package main

import "encoding/json"

// TODO: Create a CustomData type
type CustomData struct {
	Value string
}

// TODO: Implement MarshalJSON method to satisfy json.Marshaler
func (c CustomData) MarshalJSON() ([]byte, error) {
	// Return custom JSON
	return []byte(`{"custom":"` + c.Value + `"}`), nil
}

// TODO: Add a compile-time check that CustomData implements json.Marshaler
// using: var _ json.Marshaler = (*CustomData)(nil)
var _ json.Marshaler = (*CustomData)(nil)

func main() {
	data := CustomData{"test"}
	bytes, _ := json.Marshal(data)
	println(string(bytes))
}
