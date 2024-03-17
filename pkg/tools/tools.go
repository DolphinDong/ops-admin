package tools

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const (
	SecretKey = "b7eabdec-683f-0eb2-ad3c-43e3c2220251"
	Salt      = "75737eb1-3a1d-629c-6617-2c3f85769599"
)

// 生成Token：
func CreateToken(issuer string, periodMinutes int) (tokenString string, err error) {
	secretKey := []byte(SecretKey)
	m := time.Duration(periodMinutes)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * m)),
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secretKey)
	return
}

// 解析Token
func ParseToken(tokenSrt string) (claims *jwt.RegisteredClaims, err error) {
	secretKey := []byte(SecretKey)
	var token *jwt.Token
	token, err = jwt.ParseWithClaims(tokenSrt, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	claims = token.Claims.(*jwt.RegisteredClaims)
	return
}

// 生成密码
func GetEncryptedPassword(str string) string {
	str = strings.ToUpper(Salt) + str + Salt
	return MD5Str(str)
}
func MD5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
