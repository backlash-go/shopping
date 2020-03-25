package consts

const (
	CodeSuccess             = 20000 // 请求成功
	CodeNeedLogin           = 40003 //
	CodeErrParameter        = 40023
	CodeRPCError            = 40025
	CodeInternalServerError = 50000

	CodeNotAllowModify = 30000
)

var StatusText = map[int]string{
	CodeNeedLogin:           "invalid authorization",
	CodeErrParameter:        "参数错误",
	CodeInternalServerError: "服务器内部错误，请联系开发人员",
	CodeNotAllowModify:      "只能修改当前所属月份的上一月份数据",
}
