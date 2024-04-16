package handlers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type MyClaims struct {
	Uid      string `json:"user_id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var MySecret = "Uy1hjAZc2bNyUwD2kqzqn2qQpKbygVEaUNZjx3PONeQ="
var TokenExpireDuration = 4 * time.Hour
var ReTokenExpireDuration = 9 * time.Hour

// Valid 实现 jwt.Claims 接口的 Valid() 方法
func (c *MyClaims) Valid() error {
	// 检查令牌是否在过期时间之前
	if c.ExpiresAt < time.Now().Unix() {
		return errors.New("token has expired")
	}
	return nil
}

func SetToken(uid, password string) (aToken, rToken string, err error) {
	calims := MyClaims{
		Uid:      uid,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //有效时间
			IssuedAt:  time.Now().Unix(),                          //签发时间
			NotBefore: time.Now().Unix(),                          //生效时间
			Issuer:    os.Getenv("BEEN_ISSUER"),                   //签发人
			Subject:   "token",                                    //JWT ID用于标识该JWT
			Audience:  uid,                                        //用户
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &calims)
	// 生成 aToken
	aToken, err = token.SignedString([]byte(os.Getenv("MySecret")))

	// rToken 不需要存储任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ReTokenExpireDuration).Unix(), // 过期时间
		Issuer:    os.Getenv("BEEN_ISSUER"),                     // 签发人
	}).SignedString(MySecret)
	return aToken, rToken, nil
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_KEY")), nil
}

// CheckToken 解析 access_token
func CheckToken(aToken string) (claims *MyClaims, err error) {
	token, err := jwt.ParseWithClaims(aToken, &MyClaims{}, keyFunc)

	if err != nil {
		return nil, err
	}
	if !token.Valid { // token 是否有效
		err = errors.New("invalid token")
	}
	return claims, nil
}

func RefreshToken(aToken, rToken string) (newToken, newrToken string, err error) {
	// 第一步 : 判断 rToken 格式对的，没有过期的
	if _, err := jwt.Parse(rToken, keyFunc); err != nil {
		return "", "", err
	}

	// 第二步：从旧的 aToken 中解析出 claims 数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	var v *jwt.ValidationError
	_ = errors.As(err, &v)

	// 当 access token 是过期错误，并且 refresh token 没有过期就创建一个新的 access token
	if v.Errors == jwt.ValidationErrorExpired {
		return SetToken(claims.Uid, claims.Password)
	}
	return "", "", err
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		aToken := c.DefaultQuery("atoken", "")
		rToken := c.DefaultQuery("rtoken", "")

		if aToken == "" || rToken == "" {
			c.JSON(200, gin.H{
				"msg":  "没有携带 aToken 或者 rToken",
				"code": 400,
			})
			c.Abort()
			return
		}

		parasToken, err := CheckToken(aToken) // 解析 access_token
		if err == nil {                       // 当前的 access_token 格式对，没有过期
			c.JSON(200, gin.H{
				"msg":  "aToken 和 rToken 没有过期",
				"data": parasToken,
				"code": 400,
			})
			c.Next()
			return
		}

		aToken, rNewToken, err := RefreshToken(aToken, rToken)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "Need sign in again",
				"code": 400,
			})
			c.Abort()
			return
		} else {
			c.JSON(200, gin.H{
				"msg":    "aToken 和 rToken 没有过期",
				"aToken": aToken,
				"rToken": rNewToken,
				"code":   400,
			})
			c.Next()
			return
		}
	}
}