package grpc

import (
	"net/url"

	"github.com/frainzy1477/xray-core/common"
	"github.com/frainzy1477/xray-core/transport/internet"
)

const protocolName = "grpc"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}

func (c *Config) getNormalizedName() string {
	return url.PathEscape(c.ServiceName)
}
