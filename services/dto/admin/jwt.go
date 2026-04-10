package admin

import (
	"blog_backend_go/api/v1/dto/admin/auth"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// CustomClaims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID          uuid.UUID
	Account     string
	UserName    string
	Permissions auth.Permission
}
