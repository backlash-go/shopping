package handler

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"shopping/consts"
	"shopping/entity"
	"shopping/models"
	"shopping/resource"
	"shopping/service"
	"time"

	//"strconv"
)

func init() {
	routers = append(routers, handler{
		Method: consts.HttpMethodPost,
		Path:   "/login",
		Hf:     UserVerify,
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
	req := &entity.LoginRequestParam{}
	err := c.Bind(req)
	if err != nil {
		resource.GetLogger().Errorf("bin req i failed")
		return ErrorResp(c, consts.StatusText[consts.CodeLoginErrParameter], consts.CodeLoginErrParameter)
	}
	if err = c.Validate(req); err != nil {
		resource.GetLogger().Errorf("Validate req is failed")
		return ErrorResp(c, consts.StatusText[consts.CodeLoginErrParameter], consts.CodeLoginErrParameter)
	}

	user, err := service.VerifyUser(req.Account)
	if err == gorm.ErrRecordNotFound {
		return ErrorResp(c, consts.StatusText[consts.CodeAccountIsNotExist], consts.CodeAccountIsNotExist)
	}
	if err != nil {
		resource.GetLogger().Errorf("Validate req is failed:  %s", err.Error())
		return ErrorResp(c, consts.StatusText[consts.CodeInternalServerError], consts.CodeInternalServerError)
	}

	password := MdSalt(req.Password)
	if user.Password == password {
		session, err := service.CreateSession(user)
		if err != nil {
			resource.GetLogger().Errorf(" CreateSession is failed:  %s", err.Error())
			return ErrorResp(c, consts.StatusText[consts.CodeInternalServerError], consts.CodeInternalServerError)
		}
		cookie := new(http.Cookie)
		cookie.Name = consts.CookieSession
		cookie.Value = session.Key
		cookie.Expires = time.Unix(session.Expire, 0)
		cookie.Path = "/"
		c.SetCookie(cookie)
		return SuccessResp(c, user)
	}
	return ErrorResp(c, consts.StatusText[consts.CodeErrPassword], consts.CodeErrPassword)

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
		return ErrorResp(c, "注册信息不能为空", 1)
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
