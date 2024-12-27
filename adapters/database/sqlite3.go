package database

import (
	"database/sql"
	"errors"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"web1/domain/models"
)

type SQLiteDatabase struct {
	db *sql.DB
	mu sync.Mutex
}

func NewSQLiteDatabase(dataSourceName string) (*SQLiteDatabase, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Create tables if they do not exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS posts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            content TEXT
        );
        CREATE TABLE IF NOT EXISTS comments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            post_id INTEGER,
						name TEXT,
            content TEXT,
            FOREIGN KEY(post_id) REFERENCES posts(id)
        );
    `)
	if err != nil {
		return nil, err
	}

	return &SQLiteDatabase{db: db}, nil
}

func (db *SQLiteDatabase) SavePost(post *models.Post) (models.Post, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	result, err := db.db.Exec("INSERT INTO posts (name, content) VALUES (?, ?)", post.Name, post.Content)
	if err != nil {
		return models.Post{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Post{}, err
	}

	post.ID = id
	return *post, nil
}

func (db *SQLiteDatabase) ListPosts() ([]models.Post, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	rows, err := db.db.Query("SELECT id, name, content FROM posts ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Name, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (db *SQLiteDatabase) GetPost(id int64) (models.Post, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	var post models.Post
	err := db.db.QueryRow("SELECT id, name, content FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Name, &post.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Post{}, errors.New("post not found")
		}
		return models.Post{}, err
	}

	return post, nil
}

func (db *SQLiteDatabase) SaveComment(comment *models.Comment) (models.Comment, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	result, err := db.db.Exec("INSERT INTO comments (post_id, name, content) VALUES (?, ?, ?)", comment.PostID, comment.Name, comment.Content)
	if err != nil {
		return models.Comment{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	comment.ID = id
	return *comment, nil
}

func (db *SQLiteDatabase) ListComments(postID int64) ([]models.Comment, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	rows, err := db.db.Query("SELECT id, post_id, name, content FROM comments WHERE post_id = ? ORDER BY id", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.Name, &comment.Content); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (db *SQLiteDatabase) Close() error {
	return db.db.Close()
}
