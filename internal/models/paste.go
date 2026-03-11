package models

import "time"

type PasteResponse struct {
	URL string `json:"url"`
}

type PasteRequest struct {
	Content string `json:"content"`
}

type Paste struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
