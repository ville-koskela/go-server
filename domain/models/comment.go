package models

type Comment struct {
	ID 			int			`json:"id"`
	PostID	int			`json:"postId"`
	Nick		string	`json:"nick"`
	Content	string	`json:"content"`
}
