package handler

import (
	"fmt"
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




func UserVerify(c echo.Context,) error {
	req := &communication.GetLoginParam{}
	err := c.Bind(req)
	if err != nil {
		resource.GetLogger().Errorf("bin req i failed")
	}

    password := MdSalt(req.Password)

	if req.Method == "passwd" {
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
		verifyUser, err := service.VerifyPassword(req.Cellphone, password)
		if err != nil {
			resource.GetLogger().Errorf("handler job use VerifyPassword is failed %v,%v", req.Cellphone, err.Error())
		}

		if verifyUser.Cellphone != "" || verifyUser.Password != "" {
			//生成token
			token := CreateToken(verifyUser.Cellphone)
			c.Response().Header().Set("key",token)
			//存入redis
			//service.SetValue(strconv.Itoa(int(verifyUser.Id)),token,3600)
			service.HMSet(token,verifyUser.Cellphone,verifyUser.Email,3600)
			return SuccessResp(c, "登陆成功")


		}
		return ErrorResp(c, "该手机号不存在或密码错误", 1)	}

	if req.Method == "Code"{
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
	return ErrorResp(c, "请先使用密码或者验证码登陆", 1)
}

//用户注册
func Registry(c echo.Context) error {
	req := &models.User{}
	err := c.Bind(req)
	if err != nil {
		resource.GetLogger().Errorf("handler job use Registry is failed %v,%v", req.Cellphone, err.Error())
	}
	password := MdSalt(req.Password)
	user, err := service.SelectUser(req.Cellphone)
	if err != nil {
		fmt.Println(err)
	}
	if user.Cellphone != ""{
		return ErrorResp(c,"手机号已经注册",11)
	}
	err = service.RegistryUser(req.Cellphone, password, req.Email, req.Address)
	if err != nil {
		resource.GetLogger().Errorf(" Registry is failed %v,%v", req.Cellphone, err.Error())
		return ErrorResp(c, "regisry is error", 111)
	}
	return SuccessResp(c, "用户注册成功")
}



