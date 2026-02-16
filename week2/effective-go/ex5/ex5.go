package main

import "fmt"

// Hasher interface - any type that can compute a hash
type Hasher interface {
	Hash() uint32
}

// TODO: Create a SimpleHasher struct (unexported)
// with a Hash() method that returns a simple hash
type simpleHasher struct{}

func (s *simpleHasher) Hash() uint32 {
	return 12345
}

// TODO: Create a CRC32Hasher struct (unexported)
// with a Hash() method that returns a different hash
type crc32Hasher struct{}

func (s *crc32Hasher) Hash() uint32 {
	return 67890
}

// TODO: Create NewSimpleHasher() Hasher function
// that returns the interface, not the concrete type
func NewSimpleHasher() Hasher {
	return &simpleHasher{}
}

// TODO: Create NewCRC32Hasher() Hasher function
func NewCRC32Hasher() Hasher {
	return &crc32Hasher{}
}

func main() {
	h1 := NewSimpleHasher()
	h2 := NewCRC32Hasher()

	fmt.Printf("Simple: %d\n", h1.Hash())
	fmt.Printf("CRC32: %d\n", h2.Hash())

	// Show that you can swap implementations
	var h Hasher = NewSimpleHasher()
	fmt.Printf("Current: %d\n", h.Hash())
	h = NewCRC32Hasher()
	fmt.Printf("Switched: %d\n", h.Hash())
}
