package initialize

import (
	"log"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

func InitIP2Region() *xdb.Searcher {
	var err error
	dbPath := "resource/ip2region.xdb"
	//把整个库加载到内存
	buff, err := xdb.LoadContentFromFile(dbPath)
	if err != nil {
		// global.LOG.Error("加载ip2region.xdb失败", zap.Error(err))
		log.Println("加载ip2region.xdb失败", err.Error())
		return nil
	}

	// 创建基于内存的搜索器
	searcher, err := xdb.NewWithBuffer(xdb.IPv4, buff)
	if err != nil {
		log.Println("创建ip2region搜索器失败", err.Error())
		// global.LOG.Error("创建ip2region搜索器失败", zap.Error(err))
		return nil
	}
	return searcher
}
