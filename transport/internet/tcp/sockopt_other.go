// +build !linux,!freebsd

package tcp

import (
	"github.com/frainzy/xray-core/common/net"
	"github.com/frainzy/xray-core/transport/internet"
)

func GetOriginalDestination(conn internet.Connection) (net.Destination, error) {
	return net.Destination{}, nil
}
