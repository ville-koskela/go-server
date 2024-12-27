package models

type Post struct {
    ID      int64  `json:"id"`
    Nick    string `json:"nick"`
    Content string `json:"content"`
}
