package admin

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Article management
func GetArticleList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetArticleList"})
}

func AddArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddArticle"})
}

func UpdateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateArticle"})
}

func DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteArticle"})
}

func AddTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddTag"})
}

func DeleteTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteTag"})
}

// Comment management
func GetCommentList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetCommentList"})
}

func DeleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteComment"})
}

// Message management
func GetMessageList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetMessageList"})
}

func DeleteMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteMessage"})
}

// Friendship management
func GetFriendshipList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetFriendshipList"})
}

func DeleteFriendship(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteFriendship"})
}
