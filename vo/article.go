package vo

type ArticleRequest struct {
	Title   string       `json:"title,omitempty"`
	Content string       `json:"content,omitempty"`
	Tags    []TagRequest `json:"tags"`
}
