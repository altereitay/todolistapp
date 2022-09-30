package types

import (
	"github.com/google/uuid"
	"time"
)

type PostMessage struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Create string `json:"createAt"`
	Due    string `json:"due"`
}

type NoteMessage struct {
	Title  string
	Body   string
	Create time.Time
	Due    time.Time
	Id     uuid.UUID
}
