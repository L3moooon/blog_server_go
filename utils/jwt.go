package utils

import (
	"blog_backend_go/global"
	"blog_backend_go/services/dto/admin"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

// 创建jwt声明
func (j *JWT) CreateClaims(baseClaims admin.BaseClaims) (claims admin.CustomClaims) {
	bufferTime, _ := ParseDuration(global.CONFIG.JWT.BufferTime)
	expiresTime, _ := ParseDuration(global.CONFIG.JWT.ExpiresTime)
	claims = admin.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bufferTime / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"blog-backend-go"},             // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),       // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresTime)), // 过期时间 7天  配置文件
			Issuer:    global.CONFIG.JWT.Issuer,                        // 签名的发行者
		},
	}
	return claims
}

// 创建token
func (j *JWT) CreateToken(claims admin.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.CONFIG.JWT.SigningKey))
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (claims admin.CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.SigningKey), nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
		return
	}
	return
}
