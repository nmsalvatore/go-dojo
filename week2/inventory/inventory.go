package inventory

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Inventory struct {
	notifier Notifier
	books    []Book
}

func NewInventory(n Notifier) *Inventory {
	return &Inventory{
		notifier: n,
	}
}

func (inv *Inventory) Add(b Book) error {
	inv.books = append(inv.books, b)
	msg := fmt.Sprintf("%q by %s added to inventory", b.Title, b.Author)

	err := inv.notifier.Notify(msg)
	if err != nil {
		return fmt.Errorf("notifying: %w", err)
	}
	return nil
}

func (inv *Inventory) Remove(title string) error {
	for i, b := range inv.books {
		if b.Title == title {
			inv.books = append(inv.books[:i], inv.books[i+1:]...)
			msg := fmt.Sprintf("%q removed from inventory", b.Title)

			err := inv.notifier.Notify(msg)
			if err != nil {
				return fmt.Errorf("notifying: %w", err)
			}

			return nil
		}
	}
	return &BookNotFoundError{title}
}

func (inv *Inventory) List() []Book {
	result := make([]Book, len(inv.books))
	copy(result, inv.books)
	return result
}

func (inv *Inventory) Count() int {
	return len(inv.books)
}

func (inv *Inventory) SearchByTitle(query string) []Book {
	var results []Book
	for _, b := range inv.books {
		if strings.Contains(strings.ToLower(b.Title), strings.ToLower(query)) {
			results = append(results, b)
		}
	}
	return results
}

func (inv *Inventory) SearchByAuthor(author string) []Book {
	var results []Book
	for _, b := range inv.books {
		if strings.Contains(strings.ToLower(b.Author), strings.ToLower(author)) {
			results = append(results, b)
		}
	}
	return results
}

func (inv *Inventory) ExportToJSON(filename string) error {
	data, err := json.MarshalIndent(inv.books, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling inventory: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("writing file %q: %w", filename, err)
	}
	return nil
}
