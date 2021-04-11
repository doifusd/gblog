package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Descs         string `json:"descs"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
