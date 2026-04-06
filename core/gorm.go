package core

import (
	"log"

	"blog_backend_go/global"
	"blog_backend_go/model"
)

// AutoMigrate 执行所有数据库表的自动迁移
func AutoMigrate() {
	if global.DB == nil {
		log.Println("⚠️ 数据库连接未初始化，跳过自动迁移")
		return
	}

	err := global.DB.AutoMigrate(
		&model.Article{},
		&model.ArticleComment{},
		&model.ArticleLike{},
		&model.ArticleTag{},
		&model.ArticleTagMap{},
		&model.SysLogDbBackup{},
		&model.SysLogReboot{},
		&model.SysLogServer{},
		&model.SysPermission{},
		&model.SysRole{},
		&model.SysRolePermission{},
		&model.SysUser{},
		&model.SysUserRole{},
		&model.WebFriendship{},
		&model.WebMessage{},
		&model.WebVisit{},
	)
	if err != nil {
		log.Fatalf("❌ 数据库自动迁移/同步失败: %v", err)
	}
	log.Println("✅ 数据库模型自动同步成功")
}
