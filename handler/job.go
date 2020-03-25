package handler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"shopping/consts"
	"shopping/models"
	"shopping/models/communication"
	"shopping/resource"
	"shopping/service"
	//"strconv"
)

func init() {
	routers = append(routers, handler{
		Method: consts.HttpMethodPost,
		Path:   "/login",
		Hf:     UserVerify,
		//{
		//	"cellphone":"18918734197",
		//	"password":"lyqissb",
		//	"method": "passwd"
		//}
	})

	routers = append(routers, handler{
		Method: consts.HttpMethodPost,
		Path:   "/registry",
		Hf:     Registry,
	})

	routers = append(routers, handler{
		Method: consts.HttpMethodPost,
		Path:   "/loginout",
		Hf:     Registry,
	})
}

func UserVerify(c echo.Context, ) error {
	req := &communication.GetLoginParam{}
	err := c.Bind(req)
	if err != nil {
		resource.GetLogger().Errorf("bin req i failed")
		return ErrorResp(c, consts.StatusText[consts.CodeLoginErrParameter], consts.CodeLoginErrParameter)
	}
	password := MdSalt(req.Password)

	switch req.Method {
	case "passwd":
		if req.Cellphone == "" {
			return ErrorResp(c, consts.StatusText[consts.CodeInputCellphone], consts.CodeInputCellphone)
		}
		user, err := service.SelectUser(req.Cellphone)
		if err != nil {
			resource.GetLogger().Errorf("handler job use selectuser is failed %v,%v", req.Cellphone, err.Error())
			return ErrorResp(c, consts.StatusText[consts.CodePhoneIsNotExist], consts.CodePhoneIsNotExist)
		}
		if user.Cellphone == req.Cellphone && user.Password == password {
			token := CreateToken(user.Cellphone)
			c.Response().Header().Set("key", token)
			service.HMSet(token, user.Cellphone, user.Email, 3600)
			return SuccessResp(c, consts.StatusText[consts.CodeSuccess])
		}
		return ErrorResp(c, consts.StatusText[consts.CodeErrPassword], consts.CodeErrPassword)

	case "code":
		if req.Cellphone == "" {
			return ErrorResp(c, "请输入手机号", 1)
		}
		user, err := service.SelectUser(req.Cellphone)
		if err != nil {
			resource.GetLogger().Errorf("handler job use selectuser is failed %v,%v", req.Cellphone, err.Error())
		}
		if user.Cellphone == "" {
			return ErrorResp(c, "手机号不存在，请注册", 1)
		}

	}

	return ErrorResp(c, consts.StatusText[consts.CodeLoginErrMethod], consts.CodeLoginErrMethod)

}

//	if req.Method == "passwd" {
//		if req.Cellphone == "" {
//			return ErrorResp(c, "请输入手机号", 1)
//		}
//		user, err := service.SelectUser(req.Cellphone)
//		if err != nil {
//			resource.GetLogger().Errorf("handler job use selectuser is failed %v,%v", req.Cellphone, err.Error())
//			return ErrorResp(c, "手机号不存在，请注册", 1)
//		}
//		if user.Cellphone == req.Cellphone || user.Password == password {
//			token := CreateToken(user.Cellphone)
//			c.Response().Header().Set("key", token)
//			service.HMSet(token, user.Cellphone, user.Email, 3600)
//			return SuccessResp(c, consts.StatusText[consts.CodeSuccess])
//		}
//		return ErrorResp(c, "密码错误", 1)
//	}
//
//	if req.Method == "Code" {
//		if req.Cellphone == "" {
//			return ErrorResp(c, "请输入手机号", 1)
//		}
//		user, err := service.SelectUser(req.Cellphone)
//		if err != nil {
//			resource.GetLogger().Errorf("handler job use selectuser is failed %v,%v", req.Cellphone, err.Error())
//		}
//		if user.Cellphone == "" {
//			return ErrorResp(c, "手机号不存在，请注册", 1)
//		}
//	}
//	return ErrorResp(c, "请先使用密码或者验证码登陆", 1)
//}

//用户注册
func Registry(c echo.Context) error {
	req := &models.User{}
	err := c.Bind(req)

	if err != nil {
		resource.GetLogger().Errorf("handler job use Registry is failed %v,%v", req.Cellphone, err.Error())
		return ErrorResp(c, consts.StatusText[consts.CodeLoginErrParameter], consts.CodeLoginErrParameter)
	}
	if req.Cellphone == "" || req.Password == "" || req.Email == "" || req.Address == "" {
		return ErrorResp(c,"注册信息不能为空",1)
	}
	password := MdSalt(req.Password)
	_, err = service.SelectUser(req.Cellphone)
	if err != gorm.ErrRecordNotFound {
		fmt.Println(err)
		return ErrorResp(c, "手机号已经注册", 1)
	}

	err = service.RegistryUser(req.Cellphone, password, req.Email, req.Address)
	if err != nil {
		resource.GetLogger().Errorf(" Registry is failed %v,%v", req.Cellphone, err.Error())
		return ErrorResp(c, "注册失败", 111)
	}
	return SuccessResp(c, "注册成功")
}
