package utils

import (
	"log"
	"testing"
)

// 百度的ip
var baiduDsn = "202.108.22.5"
var localDsn = "127.0.0.1"

func TestIp(t *testing.T) {
	ip := GetIpSource(localDsn)

	log.Println("--->" + ip)
}
