package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleateTagFail = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
	ErrorTagExist       = NewError(20040001, "标签已经存在")

	SuccessCreateTag = NewError(100010001, "添加标签成功")
	SuccessUpdateTag = NewError(100010002, "编辑标签成功")
	SuccessDeleteTag = NewError(100010003, "删除标签成功")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	SuccessCreateArticle = NewError(100020001, "添加文章成功")
	SuccessUpdateArticle = NewError(100020002, "编辑文章成功")
	SuccessDeleteArticle = NewError(100020003, "删除文章成功")

	SuccessCreateUser = NewError(100030001, "用户注册成功")
	SuccessUpdateUser = NewError(100030002, "用户编辑成功")
	SuccessDeleteUser = NewError(100030003, "用户删除成功")
)
