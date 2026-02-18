package inventory

import "fmt"

type BookNotFoundError struct {
	Title string
}

func (e *BookNotFoundError) Error() string {
	return fmt.Sprintf("%q not found in inventory", e.Title)
}
