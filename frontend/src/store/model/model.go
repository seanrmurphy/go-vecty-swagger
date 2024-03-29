package model

import bemodels "github.com/seanrmurphy/go-vecty-swagger/models"

// Item represents a single TODO item in the store.
type Item struct {
	BackEndModel bemodels.Todo

	//ID           uuid.UUID
	//Title        string
	//Completed    bool
	//CreationDate time.Time
}

// FilterState represents a viewing filter for TODO items in the store.
type FilterState int

const (
	// All is a FilterState which shows all items.
	All FilterState = iota

	// Active is a FilterState which shows only non-completed items.
	Active

	// Completed is a FilterState which shows only completed items.
	Completed
)
