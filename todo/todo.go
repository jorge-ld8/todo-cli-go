package todo

import "time"

// Item represents a single todo item.
type Item struct {
	ID		int       `json:"id"`
	Title	string    `json:"title"`
	Completed	bool      `json:"completed"`
	CreatedAt	time.Time `json:"created_at"`
	CompletedAt	*time.Time `json:"completed_at,omitempty"`
}

// Todos represents a collection of todo items.
type Todos struct {
	Items []Item `json:"items"`
}