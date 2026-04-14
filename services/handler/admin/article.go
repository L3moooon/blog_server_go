package admin

import (
	"blog_backend_go/api/v1/dto/admin/article"
	"blog_backend_go/api/v1/dto/common/response"
	"blog_backend_go/global"
	"blog_backend_go/model"
	"fmt"
)

type ArticleService struct {
}

var ArticleServiceInstance = new(ArticleService)

// 获取文章列表
func (s *ArticleService) GetArticleList(req article.ListRequest) (*article.ListResponse, error) {
	var articles []model.Article
	var db = global.DB
	var total int64
	db = db.Model(&model.Article{})

	// 预加载标签
	db = db.Preload("Tags")
	// 关键词搜索
	if len(req.Keyword) > 0 {
		db = db.Where("a.title LIKE ?", "%"+req.Keyword+"%")
	}
	// 日期范围搜索
	if len(req.DateRange) > 0 {
		db = db.Where("a.created_at BETWEEN ? AND ?", req.DateRange[0], req.DateRange[1])
	}
	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}
	// 排序
	if len(req.Sort) > 0 {
		for _, sort := range req.Sort {
			db = db.Order(sort.Field + " " + sort.Order)
		}
	}
	// 分页
	err := db.Scopes(req.Paginate()).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(articles)
	return &article.ListResponse{
		List: articles,
		PageInfo: response.PageInfo{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
	}, nil
}

// 获取文章细节（内容）
func (s *ArticleService) GetArticleDetail(req article.DetailRequest) (*article.DetailResponse, error) {
	var article model.Article
	err := global.DB.
		Table("article a").
		Select("article.cover, article.title, article.abstract, article.content, article.view, article.like, article.status, article.top, article.created_at, article.updated_at").
		Where("id = ?", req.ID).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article.DetailResponse{
		Article: article,
	}, nil
}

// 添加文章
func (s *ArticleService) AddArticle(req article.AddRequest) error {
	var article model.Article
	article.Title = req.Title
	article.Cover = req.Cover
	article.Abstract = req.Abstract
	article.Content = req.Content
	if err := global.DB.Create(&article).Error; err != nil {
		return err
	}
	// 添加标签
	for _, tag := range req.Tags {
		if err := global.DB.Create(&model.ArticleTagMap{
			ArticleID: article.ID,
			TagID:     uint64(tag),
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

// 更新文章
func (s *ArticleService) UpdateArticle(req article.UpdateRequest) error {
	var article model.Article
	err := global.DB.Where("id = ?", req.ID).First(&article).Error
	if err != nil {
		return err
	}
	article.Title = req.Title
	article.Cover = req.Cover
	article.Abstract = req.Abstract
	article.Content = req.Content
	if err := global.DB.Save(&article).Error; err != nil {
		return err
	}
	// 更新标签
	// 先删除所有标签
	if err := global.DB.Where("article_id = ?", req.ID).Delete(&model.ArticleTagMap{}).Error; err != nil {
		return err
	}
	// 再添加新标签
	for _, tag := range req.Tags {
		if err := global.DB.Create(&model.ArticleTagMap{
			ArticleID: article.ID,
			TagID:     uint64(tag),
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

// 删除文章
func (s *ArticleService) DeleteArticle(req article.DeleteRequest) error {
	var article model.Article
	//开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() //回滚
		}
	}()
	err := global.DB.Where("id = ?", req.ID).First(&article).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := global.DB.Delete(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除标签
	if err := global.DB.Where("article_id = ?", req.ID).Delete(&model.ArticleTagMap{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
