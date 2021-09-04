package grpc

import (
	"github.com/frainzy/xray-core/common"
	"github.com/frainzy/xray-core/transport/internet"
)

const protocolName = "grpc"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
