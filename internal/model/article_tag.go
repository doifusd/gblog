package model

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleDI uint32 `json:"article_di"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
