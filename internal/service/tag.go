package service

type CountTagRequest struct {
	Name  string `form:"name" bindding:"max=100"`
	State uint8  `form:"state,default=1" bindding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" bindding:"max=100"`
	State uint8  `form:"state,default=1" bindding:"oneof=0 1"`
}

type CrateTagResquest struct {
	Name      string `form:"name" bindding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" bindding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" bindding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" bindding:"required,gte=1"`
	Name       string `form:"name" bindding:"min=3,max=100"`
	State      uint8  `form:"state,default=1" bindding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" bindding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" bindding:"required,gte=1"`
}
