package core

import (
	"log"

	"blog_backend_go/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDataBase() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.DbName == "" {
		log.Println("⚠️ 数据库名称为空，跳过数据库自动连接")
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	// 初始化 GORM
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 迁移时禁用外键约束，以防出现循环引用错误
	})

	if err != nil {
		log.Printf("❌ MySQL 连接失败: %v", err)
		return nil
	}

	// 设置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)

	log.Println("✅ 数据库成功连接")
	return db
}
