package errcode

var (
	SuccessCreateUser     = NewError(101200, "用户注册成功")
	SuccessLoginUser      = NewError(102200, "用户登录成功")
	SuccessUpdateUser     = NewError(103200, "用户编辑成功")
	SuccessDeleteUser     = NewError(104200, "用户删除成功")
	ErrorUserExist        = NewError(101201, "用户已经存在")
	ErrorCreateUserFail   = NewError(101400, "用户注册失败")
	ErrorUserNotExist     = NewError(100404, "用户不存在")
	ErrorUserPasswordFail = NewError(100412, "用户密码错误")

	SuccessCreateTag    = NewError(200200, "添加标签成功")
	SuccessUpdateTag    = NewError(200201, "编辑标签成功")
	SuccessDeleteTag    = NewError(200202, "删除标签成功")
	SuccessGetTag       = NewError(200203, "获取标签成功")
	ErrorCreateTagFail  = NewError(200501, "创建标签失败")
	ErrorUpdateTagFail  = NewError(200502, "更新标签失败")
	ErrorDeleateTagFail = NewError(200503, "删除标签失败")
	ErrorGetTagListFail = NewError(200504, "获取标签列表失败")
	ErrorCountTagFail   = NewError(200505, "统计标签失败")
	ErrorTagExist       = NewError(200206, "标签已经存在")
	ErrorGetTagFail     = NewError(200507, "获取标签失败")

	SuccessCreateArticle    = NewError(300200, "添加文章成功")
	SuccessUpdateArticle    = NewError(300201, "编辑文章成功")
	SuccessDeleteArticle    = NewError(300202, "删除文章成功")
	SuccessGetArticle       = NewError(300203, "获取文章成功")
	ErrorCreateArticleFail  = NewError(300501, "创建文章失败")
	ErrorUpdateArticleFail  = NewError(300502, "更新文章失败")
	ErrorDeleateArticleFail = NewError(300503, "删除文章失败")
	ErrorGetArticleListFail = NewError(300504, "获取文章列表失败")
	ErrorCountArticleFail   = NewError(300505, "统计文章失败")
	ErrorGetArticleFail     = NewError(300506, "获取文章失败")

	ErrorUploadFileFail = NewError(600500, "上传文件失败")
)
