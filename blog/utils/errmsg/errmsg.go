package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code=1000...用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_FORMAT   = 1007
	ERROR_USER_NO_RIGHT  = 1008

	// code=2000...分类模块的错误
	ERROR_CATENAME_USED = 2001
	ERROR_CAT_NOT_EXIST = 2002

	// code=3000...文章模块的错误
	ERROR_ART_NOT_EXIST = 3001

	// code=4000... login包错误
	ERROR_TOKEN_GEN = 40001

	// code= 5000... Tag模块错误
	ERROR_TAG_NOT_EXIST = 50001
	ERROR_TAG_EXIST     = 50002
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST:    "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:  "TOKEN已过期",
	ERROR_TOKEN_WRONG:    "TOKEN不正确",
	ERROR_TOKEN_FORMAT:   "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:  "该用户无权限登录后台",

	ERROR_CATENAME_USED: "分类名已存在",
	ERROR_CAT_NOT_EXIST: "该分类不存在",

	ERROR_ART_NOT_EXIST: "文章不存在",
	ERROR_TOKEN_GEN:     "token生成错误",

	ERROR_TAG_NOT_EXIST: "标签不存在",
	ERROR_TAG_EXIST:     "标签已经存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
