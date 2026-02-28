package uip

import (
	"fmt"
	"strings"

	"github.com/chasespace/goutil/uerr"
	"github.com/lionsoul2014/ip2region/binding/golang/service"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type ip2RegionSearcher interface {
	SearchByStr(string) (string, error)
	Close()
}

var Ip2Region ip2RegionSearcher

type Ip2RegionDetail struct {
	Country    string // 中国 美国
	CountryISO string // CN US
	Province   string // 广东
	City       string // 深圳
}

func replaceChars(chars string) string {
	if chars == "0" {
		return ""
	}
	return chars
}

func SearchIPv4(ip string) (*Ip2RegionDetail, error) {
	if ip == "" {
		return nil, uerr.NewWithDetail(uerr.CodeBadRequest, "ip is empty")
	}
	s, err := Ip2Region.SearchByStr(ip)
	if err != nil {
		return nil, uerr.NewWithError(uerr.CodeInternalError, err, "Failed to search ip")
	}
	ss := strings.Split(s, "|")
	if len(ss) >= 5 {
		if strings.Contains(ss[1], "香港") || strings.Contains(ss[1], "澳门") || strings.Contains(ss[1], "台湾") {
			ss[1] = string([]rune(ss[1])[:2])
		}
		return &Ip2RegionDetail{
			CountryISO: replaceChars(ss[4]),
			Country:    replaceChars(ss[0]),
			Province:   strings.TrimRight(replaceChars(ss[1]), "省"),
			City:       strings.TrimRight(replaceChars(ss[2]), "市"),
		}, nil
	}
	return nil, uerr.NewWithError(uerr.CodeInternalError, err, "Failed to parse ip")
}

func MustInitIp2Region(Ip2regionXdbPath string) {
	//验证 xdb 文件的适用性
	err := xdb.VerifyFromFile(Ip2regionXdbPath)
	if err != nil {
		panic(fmt.Errorf("failed to verify ip2region xdb file: %s", err))
	}

	// 1, 创建 v4 的配置：指定缓存策略和 v4 的 xdb 文件路径
	// 参数1： 缓存策略, options: service.NoCache / service.VIndexCache / service.BufferCache
	// 参数2: xdb 文件路径
	// 参数3: 初始化的查询器数量
	v4Config, err := service.NewV4Config(service.VIndexCache, Ip2regionXdbPath, 10)
	if err != nil {
		panic(fmt.Errorf("failed to create v4 config: %s", err))
	}

	// 2，通过上述配置创建 Ip2Region 查询服务
	Ip2Region, err = service.NewIp2Region(v4Config, nil)
	if err != nil {
		panic(fmt.Errorf("failed to create ip2region service: %s", err))
	}

	// 3，导出 ip2region 服务进行双版本的IP地址的并发查询，例如：
	//v4, err := SearchIPv4("8.8.8.8") // 进行 IPv4 查询
	//if err != nil {
	//	panic(fmt.Errorf("failed to search ip: %s", err))
	//}
	//fmt.Printf("test SearchIPv4: %+v\n", v4) // 中国|广东省|深圳市|电信|CN
}

func CloseIp2Region() {
	if Ip2Region != nil {
		Ip2Region.Close()
	}
}
