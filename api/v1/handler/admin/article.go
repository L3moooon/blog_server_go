package admin

import (
	"blog_backend_go/api/v1/dto/admin/article"
	"blog_backend_go/api/v1/dto/common/response"
	"blog_backend_go/utils"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

func (s *ArticleApi) GetArticleList(c *gin.Context) {
	var req article.ListRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	res, err := articleService.GetArticleList(req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetailed(c, res, "获取文章成功")
}
func (s *ArticleApi) AddArticle(c *gin.Context) {
	var req article.AddRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"Title":    {utils.NotEmpty()}, // 标题不能为空
		"Cover":    {utils.NotEmpty()}, // 封面不能为空
		"Abstract": {utils.NotEmpty()}, // 摘要不能为空
		"Content":  {utils.NotEmpty()}, // 内容不能为空
		"Tags":     {utils.NotEmpty()}, // 标签不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = articleService.AddArticle(req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetailed(c, nil, "添加文章成功")
}
func (s *ArticleApi) UpdateArticle(c *gin.Context) {
	var req article.UpdateRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"ID": {utils.NotEmpty()}, // ID不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = articleService.UpdateArticle(req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetailed(c, nil, "更新文章成功")
}

func (s *ArticleApi) DeleteArticle(c *gin.Context) {
	var req article.DeleteRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// 字段校验
	var verifyRule = utils.Rules{
		"ID": {utils.NotEmpty()}, // ID不能为空
	}
	err = utils.Verify(req, verifyRule)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	err = articleService.DeleteArticle(req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetailed(c, nil, "删除文章成功")
}
