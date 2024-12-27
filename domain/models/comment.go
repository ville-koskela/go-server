package models

type Comment struct {
	ID      int64  `json:"id"`
	PostID  int64  `json:"postId"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
