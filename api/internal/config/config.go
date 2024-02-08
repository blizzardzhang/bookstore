package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf //增加add服务依赖
	Check zrpc.RpcClientConf //增加check服务依赖
}
