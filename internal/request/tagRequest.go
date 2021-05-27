package request

type CountTagRequest struct {
	Name  string `json:"name" binding:"omitempty,max=100"`
	State uint8  `json:"state,default=1" binding:"omitempty,oneof=0 1"`
}

type TagInfoRequest struct {
	TagId uint32 `json:"tag_id" binding:"required"`
}
type TagListRequest struct {
	Name  string `json:"name" binding:"omitempty,max=100"`
	State uint8  `json:"state" binding:"omitempty,oneof=0 1"`
}

type CreateTagResquest struct {
	Name      string `json:"name" binding:"required,min=3,max=100"`
	CreatedBy uint32 `json:"created_by"`
}

type UpdateTagRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
	Name       string `json:"name" binding:"min=1,max=100"`
	State      uint8  `json:"state,default=1" binding:"required,oneof=0 1"`
	ModifiedBy uint32 `json:"modified_by"`
}

type DeleteTagRequest struct {
	ID         uint32 `json:"id" binding:"required,gte=1"`
	ModifiedBy uint32 `json:"modified_by"`
}
