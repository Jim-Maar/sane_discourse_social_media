package types

type PostType string

const (
	PostTypeArticle PostType = "article"
	PostTypeBlog    PostType = "blog"
	PostTypeVideo   PostType = "video"
	PostTypePodcast PostType = "podcast"
	PostTypePaper   PostType = "paper"
)
