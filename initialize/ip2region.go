package initialize

import (
	"log"

	"github.com/lionsoul2014/ip2region/binding/golang/service"
)

func InitIP2Region() *service.Ip2Region {
	v4Config, err := service.NewV4Config(service.VIndexCache, "ip2region v4 xdb path", 20)
	if err != nil {
		log.Println("failed to create v4 config", err)
		return nil
	}

	v6Config, err := service.NewV6Config(service.VIndexCache, "ip2region v6 xdb path", 20)
	if err != nil {
		log.Println("failed to create v6 config", err)
		return nil
	}

	ip2region, err := service.NewIp2Region(v4Config, v6Config)
	if err != nil {
		log.Println("failed to create ip2region service", err)
		return nil
	}

	// v4Region, err := ip2region.Search("113.92.157.29")                          // 进行 IPv4 查询
	// v6Region, err := ip2region.Search("240e:3b7:3272:d8d0:db09:c067:8d59:539e") // 进行 IPv6 查询

	// ip2region.Close()
	return ip2region
}
