package models

import "time"

type SMS struct {
	Id        string    `json:"id"`
	Number    string    `json:"number"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
