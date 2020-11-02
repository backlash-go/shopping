package consts

const (
	CodeSuccess              	= 20000 // 登录成功
	CodeLoginErrParameter    	= 40003 //
	CodeInternalServerError  	= 50000
	CodeAccountIsNotExist         = 60000
	CodeErrPassword  			= 60010
	CodeInputCellphone			= 60011
	CodeLoginErrMethod				= 60020
)

var StatusText = map[int]string{
	CodeSuccess:			 	"登录成功",
	CodeLoginErrParameter:   	"登录参数错误",
	CodeInternalServerError: 	"服务器内部错误，请联系开发人员",
	CodeAccountIsNotExist: 		"账号不存在,请注册",
	CodeErrPassword:			"密码错误",
	CodeInputCellphone:			"请输入手机号",
	CodeLoginErrMethod:			"Method Error Parameter",
}
