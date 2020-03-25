package middle

import (
	"github.com/labstack/echo/v4"
	"shopping/handler"
	"shopping/resource"
	"shopping/service"
	"strings"
)

func LoginValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if strings.HasPrefix(c.Request().URL.Path,"/login"){
			return next(c)
		}
		if strings.HasPrefix(c.Request().URL.Path,"/registry"){
			return next(c)
		}

		authorization :=  c.Request().Header.Get("key")
		if authorization == ""{
			return handler.ErrorResp(c,"header token 不存在",11)
		}

		Exists,err := service.ExistKey(authorization)
		if err != nil {
			resource.GetLogger().Errorf("ExistKEY happen %s",err)
			return handler.ErrorResp(c,"查询KEY 错误",1)
		}
		if Exists == 1{
			return next(c)
		}

		if Exists == 0{
			return handler.ErrorResp(c,"token 过期",1)
		}
		return handler.ErrorResp(c,"您需要重新登录",11)
	}

}
