package udp

import (
	"github.com/frainzy1477/xray-core/common"
	"github.com/frainzy1477/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
