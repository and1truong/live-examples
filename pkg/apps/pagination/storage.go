package pagination

import (
	"fmt"
)

var items []string

func init() {
	for i := 0; i < 100; i++ {
		items = append(items, fmt.Sprintf("This is item %d", i))
	}
}

func queryItems(page int) []string {
	start := page * itemsPerPage
	end := start + itemsPerPage
	if start >= len(items) || end > len(items) {
		return []string{}
	}

	return items[start:end]
}
