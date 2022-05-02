package middleware

import (
	"blog/utils"
	"blog/utils/errmsg"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	ss, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return ss, errmsg.SUCCESS
}

// 验证token
func checkToken(tokenStr string) (*MyClaims, int) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		fmt.Printf("ParseWithClaims解析失败", err)
	}
	if claims, _ := token.Claims.(*MyClaims); token.Valid {
		return claims, errmsg.SUCCESS
	} else {
		fmt.Println("******", claims)
		return nil, errmsg.ERROR
	}
}

// 声明jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log.Println("中间件")
		// 获取请求头部的Authorization字段进行解析得到token
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		// 错误一：
		if tokenHeader == "" {
			code = errmsg.ERROR
			c.JSON(200, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			// Abort之后不会执行剩下的所有中间件了
			// 因为可能后续有很多的中间件需要进行执行，Abort表示如果当前失败的话，则下面的中间件不会被执行
			c.Abort()
			return
		}
		checkedToken := strings.SplitN(tokenHeader, " ", 2)
		// 错误二：
		if len(checkedToken) != 2 && checkedToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(200, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// 验证token
		key, tCode := checkToken(checkedToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(200, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(200, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		// Next表示继续执行后续的中间件
		c.Next()
	}
}
