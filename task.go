package main

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"` // "petit", "moyen", "grand"
	Completed   bool      `json:"completed"`
	DueDate     string    `json:"due_date"` // Format: AAAA-MM-JJ
	CreatedAt   time.Time `json:"created_at"`
}
