package Post

type Post struct {
	Title string
	Content string
}

func NewPost(title, content string) *Post {
	return &Post{title, content}
}
