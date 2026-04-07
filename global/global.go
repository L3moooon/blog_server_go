package global

import (
	"blog_backend_go/config"

	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/v3/client"
	"github.com/gin-gonic/gin"
	"github.com/lionsoul2014/ip2region/binding/golang/service"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	LOG    *zap.Logger

	ROUTERS gin.RoutesInfo
	// ACTIVE_DBNAME *string
	// MCP_SERVER    *server.MCPServer
	BlackCache local_cache.Cache //jwt黑名单
	// lock          sync.RWMutex

	IP2REGION *service.Ip2Region
	Mailer    *gomail.Dialer
	SMS       *dypnsapi.Client
)
