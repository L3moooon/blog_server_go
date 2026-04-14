package article

import (
	"blog_backend_go/api/v1/dto/common/response"
	"blog_backend_go/model"
)

// 文章列表响应
type ListResponse struct {
	List []model.Article `json:"list"`
	response.PageInfo
}
