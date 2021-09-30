package Utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	//盐
	secret                 = []byte("BugQian")      // 后续加密增加盐增加复杂度
	TokenExpired     error = errors.New("Token已过期") // token错误类型提炼
	TokenNotValidYet error = errors.New("Token未生效") // token错误类型提炼
	TokenMalformed   error = errors.New("Token错误")  // token错误类型提炼
	TokenInvalid     error = errors.New("非法Token")  // token错误类型提炼
)

type JwtClaims struct {
	*jwt.StandardClaims
	Uid      int    `json:"uid"`
	UserName string `json:"username"`
}

//生成Token
func CreateJwtToken(uid int, UserName string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 6).Unix()
	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(expireTime),
			Issuer:    "BugQian",
		},
		uid,
		UserName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//解析Token
func ParseJwtToken(jwtToken string) (*JwtClaims, error) {
	var jwtClaim = &JwtClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, jwtClaim, func(*jwt.Token) (interface{}, error) {
		//得到盐
		return secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	return jwtClaim, nil
}
