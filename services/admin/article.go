package admin

import (
	dto "blog_backend_go/internal/dto/admin"
	model "blog_backend_go/internal/model"
)

type ArticleService struct {
}

var ArticleServiceInstance = new(ArticleService)

func (s *ArticleService) CreateArticle(u dto.LoginRequest) (userInter model.SysUser, err error) {
	print("111")
	return
}
