package request

type CountArticleRequest struct {
	CreatedBy string `form:"created_by" binding:"max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	CreatedBy string `form:"created_by" binding:"max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CrateArticleResquest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"descs" binding:"required,min=3,max=300"`
	Content       string `form:"content" binding:"required,min=3,max=10000000"`
	CoverImageUrl string `form:"img" binding:"min=3,max=1000"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	State      uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
