package middle

import (
	"github.com/labstack/echo/v4"
	"shopping/consts"
	"shopping/handler"
	"shopping/resource"
)

var (
	SkipKeyword = []string{"/health", "login"}
	defaultSkip = func(c echo.Context) bool {
		for _, word := range SkipKeyword {
			if c.Request().URL.Path == word {
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
		result, err := resource.RedisHmGet(token, consts.AuthId, consts.AuthAccount, consts.AuthCellphone, consts.AuthNickName, consts.AuthRealName, consts.AuthAvatarUrl)
		s := make([]string, len(result))
		for i, v := range result {
			s[i] = v.(string)
		}
		if len(s) ==0 {
			resource.GetLogger().Errorf(" resource.RedisHmGet(token failed:  %s", err.Error())
			return handler.ErrorResp(c, consts.StatusText[consts.CodeNeedLogin], consts.CodeNeedLogin)
		}
		if err != nil {
			resource.GetLogger().Errorf(" resource.RedisHmGet(token failed:  %s", err.Error())
			return handler.ErrorResp(c, consts.StatusText[consts.CodeInternalServerError], consts.CodeInternalServerError)
		}
		c.Request().Header.Set(consts.AuthId, s[0])
		c.Request().Header.Set(consts.AuthAccount, s[1])
		c.Request().Header.Set(consts.AuthCellphone, s[2])
		c.Request().Header.Set(consts.AuthNickName, s[3])
		c.Request().Header.Set(consts.AuthRealName, s[4])
		c.Request().Header.Set(consts.AuthAvatarUrl, s[5])

		return next(c)
	}
}
