package inventory

import (
	"errors"
	"fmt"
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		books []Book
	}{
		{
			name: "one book",
			books: []Book{
				{Title: "Learning Go", Author: "Jon Bodner"},
			},
		},
		{
			name: "two books",
			books: []Book{
				{Title: "E. M. Cioran", Author: "All Gall is Divided"},
				{Title: "Klara and the Sun", Author: "Kazuo Ishiguro"},
			},
		},
		{
			name:  "no books",
			books: []Book{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notifier := StoreInMemory{}
			inventory := NewInventory(&notifier)

			for _, book := range tt.books {
				err := inventory.Add(book)
				if err != nil {
					t.Fatal(err)
				}
			}

			got := notifier.Messages()
			for _, book := range tt.books {
				want := fmt.Sprintf("%q by %s added to inventory", book.Title, book.Author)

				if !slices.Contains(got, want) {
					t.Errorf("got %q, want %q", got, want)
				}
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name   string
		titles []string
	}{
		{
			name:   "one book",
			titles: []string{"Tears and Saints"},
		},
		{
			name:   "two books",
			titles: []string{"Tears and Saints", "A Short History of Decay"},
		},
		{
			name:   "no books",
			titles: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notifier := StoreInMemory{}
			inventory := NewInventory(&notifier)

			setupInventory(inventory)
			notifier.Reset()
			startCount := inventory.Count()

			for _, title := range tt.titles {
				err := inventory.Remove(title)
				if err != nil {
					t.Fatal(err)
				}
			}

			finalCount := inventory.Count()
			wantCount := startCount - len(tt.titles)
			if finalCount != wantCount {
				t.Errorf("got count %d, want %d", finalCount, wantCount)
			}

			messages := notifier.Messages()
			for _, title := range tt.titles {
				message := fmt.Sprintf("%q removed from inventory", title)

				if !slices.Contains(messages, message) {
					t.Errorf("%q not contained in %v", message, messages)
				}
			}
		})
	}
}

func TestRemoveError(t *testing.T) {
	notifier := StoreInMemory{}
	inventory := NewInventory(&notifier)
	setupInventory(inventory)

	err := inventory.Remove("The Bible")
	if err == nil {
		t.Error("wanted error, didn't get one")
	}

	var bnfe *BookNotFoundError
	if !errors.As(err, &bnfe) {
		t.Error("wanted BookNotFoundError, didn't get one")
	}
}

func setupInventory(inv *Inventory) {
	books := []Book{
		{
			Title:  "On the Heights of Despair",
			Author: "E. M. Cioran",
		},
		{
			Title:  "Tears and Saints",
			Author: "E. M. Cioran",
		},
		{
			Title:  "A Short History of Decay",
			Author: "E. M. Cioran",
		},
		{
			Title:  "The Temptation to Exist",
			Author: "E. M. Cioran",
		},
	}

	for _, book := range books {
		inv.Add(book)
	}
}
