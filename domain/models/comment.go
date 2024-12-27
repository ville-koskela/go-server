package models

type Comment struct {
	ID      int64  `json:"id"`
	PostID  int64  `json:"postId"`
	Nick    string `json:"nick"`
	Content string `json:"content"`
}
