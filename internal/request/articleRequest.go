package request

type CountArticleRequest struct {
	CreatedBy uint32 `json:"created_by" binding:"required"`
	State     uint8  `json:"state,default=1" binding:"oneof=0 1"`
}

type ArticleInfoRequest struct {
	ArticleID uint32 `json:"article_id" binding:"required"`
}

type ArticleListRequest struct {
	CreatedBy uint32 `json:"created_by" binding:"required"`
	State     uint8  `json:"state" binding:"omitempty,oneof=1 2"`
}

type CreateArticleResquest struct {
	Title     string   `json:"title" binding:"required,min=3,max=100"`
	Desc      string   `json:"desc" binding:"required,min=3,max=300"`
	Content   string   `json:"content" binding:"required,min=3"`
	Cover     string   `json:"cover" binding:"min=3"`
	CreatedBy uint32   `json:"created_by"`
	State     uint8    `json:"state,default=1"`
	Tags      []uint32 `json:"tags"`
}

type UpdateArticleRequest struct {
	ID         uint32   `json:"id" binding:"required,gte=1"`
	Title      string   `json:"title" binding:"required,min=3,max=100"`
	Desc       string   `json:"desc" binding:"required,min=3,max=300"`
	Content    string   `json:"content" binding:"required,min=3"`
	Cover      string   `json:"cover"`
	State      uint8    `json:"state" binding:"omitempty,oneof=1 2"`
	ModifiedBy uint32   `json:"modified_by"`
	Tags       []uint32 `json:"tags"`
}

type DeleteArticleRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
	ModifiedBy uint32 `json:"modified_by"`
}
