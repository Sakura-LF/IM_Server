package etcd

import (
	"IM_Server/core"
	"IM_Server/utils/ip"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

func PutService(etcdAddr string, serviceName string, addr string) {
	list := strings.Split(addr, ":")
	if len(list) != 2 {
		logx.Errorf("地址错误：%s", addr)
		return
	}
	if list[0] == "0.0.0.0" {
		getIP, err := ip.GetIP()
		if err != nil {
			logx.Errorf("IP获取失败: %s", err.Error())
			return
		}
		addr = strings.ReplaceAll(addr, "0.0.0.0", getIP)
	}
	client := core.InitEtcd(etcdAddr)
	_, err := client.Put(context.Background(), serviceName, addr)
	if err != nil {
		logx.Errorf("地址上传失败：%s", err.Error())
		return
	}
	logx.Infow("地址上传成功", logx.Field("serviceName", serviceName), logx.Field("Addr", addr))
}

func GetService(etcdAddr string, serviceName string) (addr string) {
	client := core.InitEtcd(etcdAddr)
	res, err := client.Get(context.Background(), serviceName)
	if err == nil && len(res.Kvs) > 0 {
		return string(res.Kvs[0].Value)
	}
	return ""
}
