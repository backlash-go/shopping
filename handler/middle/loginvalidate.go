package middle

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"shopping/handler"
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
			return handler.ErrorResp(c,"token 不存在",11)
		}

		Exists,err := service.ExistKey(authorization)
		if err != nil {
			fmt.Println(err)
		}
		if Exists == 0{
			return handler.ErrorResp(c,"token 过期",11)

		}
		return next(c)
	}

}
