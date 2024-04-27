package domain

import "time"

type ReplyID int64

type Reply struct {
	ReplyID  ReplyID   `json:"reply_id" db:"reply_id"`
	PostID   PostID    `json:"post_id" db:"post_id"`
	UserID   UserID    `json:"user_id" db:"user_id"`
	Title    string    `json:"title" db:"title"`
	Comment  string    `json:"comment" db:"comment"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}
