package models

import "time"

type Comment struct {
	ID            string    `db:"id" json:"id"`
	Comment       string    `db:"comment" json:"comment"`
	CommenterName string    `db:"commenter_name" json:"commenter_name"`
	CreatedAt     time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time `db:"updatedAt" json:"updatedAt"`
	PostID        string    `db:"post_id" json:"post_id"`
	Post          Post      `db:"-" json:"post"`
}
