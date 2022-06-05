package e

const (
	SUCCESS       = 0
	ERROR         = 500
	InvalidParams = 400
	UNDOSUCCESS   = 1

	//成员错误
	ErrorExistUser      = 10002
	ErrorNotExistUser   = 10003
	ErrorFailEncryption = 10006
	ErrorNotCompare     = 10007

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthToken             = 30003
	ErrorAuth                  = 30004
	ErrorDatabase              = 40001

	SuccessUpLoadFile   = 20000
	ErrorUpLoadFile     = 20001
	OutOfUserPermission = 20002
)
