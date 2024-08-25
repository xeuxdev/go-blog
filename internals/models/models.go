package models

import "time"

type User struct {
	ID       string `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Posts    []Post `db:"-" json:"posts"`
}

type Post struct {
	ID          string    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	FullContent string    `db:"full_content" json:"full_content"`
	Image       string    `db:"image" json:"image"`
	LikesCount  int       `db:"likesCount" json:"likesCount"`
	ViewCount   int       `db:"viewCount" json:"viewCount"`
	CreatedAt   time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt" json:"updatedAt"`
	UserID      string    `db:"user_id" json:"user_id"`
	Comments    []Comment `db:"-" json:"comments"`
	Author      User      `db:"-" json:"author"`
}

type Comment struct {
	ID            string    `db:"id" json:"id"`
	Comment       string    `db:"comment" json:"comment"`
	CommenterName string    `db:"commenter_name" json:"commenter_name"`
	CreatedAt     time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time `db:"updatedAt" json:"updatedAt"`
	PostID        string    `db:"post_id" json:"post_id"`
	Post          Post      `db:"-" json:"post"`
}
