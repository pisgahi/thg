package mdl

import "time"

type Card struct {
	PostId    string    `json:"postId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	TimeStamp string    `json:"date"`
	CreatedAt time.Time `json:"createdAt" `
}
