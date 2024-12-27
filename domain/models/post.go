package models

type Post struct {
    ID      int    `json:"id"`
    Nick    string `json:"nick"`
    Content string `json:"content"`
}
