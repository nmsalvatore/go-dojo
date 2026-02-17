package inventory

type Genre string

const (
	Fiction    Genre = "fiction"
	NonFiction Genre = "nonfiction"
	Science    Genre = "science"
	History    Genre = "history"
)

type Book struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Genre  Genre   `json:"genre"`
	Price  float64 `json:"price"`
}
