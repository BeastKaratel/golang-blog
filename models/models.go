package models

//Post is model Post
type Post struct {
	Id      string
	Title   string
	Content string
}

// NewPost is contucter of model Post
func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}
