package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf //服务发现配置
	Check zrpc.RpcClientConf
}
