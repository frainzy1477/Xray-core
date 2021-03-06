package all

import (
	"github.com/frainzy1477/xray-core/main/commands/all/api"
	"github.com/frainzy1477/xray-core/main/commands/all/tls"
	"github.com/frainzy1477/xray-core/main/commands/base"
)

// go:generate go run github.com/frainzy1477/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		//cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
