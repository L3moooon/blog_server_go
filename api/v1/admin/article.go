package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleApi struct {
}

func (s *ArticleApi) GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetArticleList"})
}
