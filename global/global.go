package global

import (
	"blog_backend_go/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  redis.UniversalClient
	CONFIG config.Server
	VP     *viper.Viper
	LOG    *zap.Logger

	ROUTERS gin.RoutesInfo
	// ACTIVE_DBNAME *string
	// MCP_SERVER    *server.MCPServer
	BlackCache local_cache.Cache //jwt黑名单
	// lock          sync.RWMutex
)
