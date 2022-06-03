package e

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	ERROR:         "fail",
	InvalidParams: "请求参数错误",
	UNDOSUCCESS:   "调用成功 操作取消或无需操作",

	ErrorExistUser:    "此用户名已存在",
	ErrorNotExistUser: "不存在此账号",
	ErrorNotCompare:   "密码错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorDatabase:              "数据库操作出错,请重试",

	SuccessUpLoadFile:   "文件上传成功",
	ErrorUpLoadFile:     "文件上传失败",
	OutOfUserPermission: "用户越权操作",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
