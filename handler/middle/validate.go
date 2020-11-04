package middle

import (
	"github.com/labstack/echo/v4"
	"log"
	"shopping/consts"
	"shopping/handler"
	"shopping/resource"
)

var (
	SkipKeyword = []string{"/health", "/login"}
	defaultSkip = func(c echo.Context) bool {
		for _, word := range SkipKeyword {
			if c.Request().URL.Path == word {
				log.Println(c.Request().URL.Path)
				return true
			}
		}
		return false
	}
)

// ValidateToken token验证
func Valid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if defaultSkip(c) {
			return next(c)
		}
		cookie, err := c.Cookie(consts.CookieSession)
		if err != nil {
			return err
		}
		token := cookie.Value
		log.Println(token)
		result, err := resource.RedisHmGet(token, consts.AuthId, consts.AuthAccount, consts.AuthCellphone, consts.AuthNickName, consts.AuthRealName, consts.AuthAvatarUrl)
		if err != nil {
			resource.Logger.Errorf(" resource.RedisHmGet(token failed:  %s", err.Error())
			return handler.ErrorResp(c, consts.StatusText[consts.CodeInternalServerError], consts.CodeInternalServerError)
		}
		s := make([]string, 0)
		for i, v := range result {
			t, ok := v.(string)
			if ok {
				s[i] = t
			}
		}
		if len(s) == 0 {
			resource.Logger.Errorf(" resource.RedisHmGet(token failed")
			return handler.ErrorResp(c, consts.StatusText[consts.CodeNeedLogin], consts.CodeNeedLogin)
		}
		return next(c)
	}
}
