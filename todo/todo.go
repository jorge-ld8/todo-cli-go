package todo

import (
	"encoding/json"
	"os"
	"time"
)

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


// add todo item
func (t *Todos) Add(title string) Item {
	newID := 1
	if len(t.Items) > 0 {
		newID = t.Items[len(t.Items)-1].ID + 1
	}

	// create the item
	item := Item{
		ID:		newID,
		Title:	title,
		Completed:	false,
		CreatedAt:	time.Now(),
	}

	// append the item
	t.Items = append(t.Items, item)

	// return the item
	return item
}

// get list of todo items
func (t *Todos) List() []Item {
	return t.Items
}

// FILE PERSISTENCE methods
const defaultFile string = ".todos.json"

// save a file
func (t *Todos) Save() error {
	data, err := json.MarshalIndent(t, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(defaultFile, data, 0644)
}

// load from file
func (t *Todos) Load() error {
	data, err := os.ReadFile(defaultFile)
	if err != nil {
		if os.IsNotExist(err) {
			// No existing file, start with empty todos
			t.Items = []Item{}
			return nil
		}
		return err
	}
	
	return json.Unmarshal(data, t)
}