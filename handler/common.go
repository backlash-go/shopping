package handler

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"shopping/consts"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var routers []handler

type handler struct {
	Hf           echo.HandlerFunc
	Method, Path string
}

func GetRouters() []handler {
	return routers
}

func SuccessResp(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "ok",
		"code": consts.CodeSuccess,
		"data": data,

	})
}

func ErrorResp(c echo.Context, msg string, code int) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  msg,
		"code": code,
	})
}

func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "Internal Server Error",
		"code": "50000",
		"data": nil,
	})
}


//密码加密加盐
func MdSalt(p string) string{
	salt := []byte("$%*&%99")
	hashmd := md5.New()
	io.WriteString(hashmd,p)
	password := fmt.Sprintf("%x", hashmd.Sum(salt))
	return password
}

//生成token
func CreateToken(cellphone string) string {
	h := sha256.New()
	rand.Seed(time.Now().UnixNano())
	io.WriteString(h, cellphone)
	io.WriteString(h, strconv.Itoa(rand.Int()))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}