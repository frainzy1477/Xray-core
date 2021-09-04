package command_test

import (
	"context"
	"testing"

	"github.com/frainzy/xray-core/app/dispatcher"
	"github.com/frainzy/xray-core/app/log"
	. "github.com/frainzy/xray-core/app/log/command"
	"github.com/frainzy/xray-core/app/proxyman"
	_ "github.com/frainzy/xray-core/app/proxyman/inbound"
	_ "github.com/frainzy/xray-core/app/proxyman/outbound"
	"github.com/frainzy/xray-core/common"
	"github.com/frainzy/xray-core/common/serial"
	"github.com/frainzy/xray-core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
