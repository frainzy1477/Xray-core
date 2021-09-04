// +build !linux,!freebsd

package tcp

import (
	"github.com/frainzy1477/xray-core/common/net"
	"github.com/frainzy1477/xray-core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
