package request

type LoginRequest struct {
	Mobile string `from:"mobile" binding:"required,min=11,max=11"`
	Passwd string `form:"passwd" binding:"required"`
}
