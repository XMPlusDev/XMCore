package all

import (
	"github.com/xcode75/xcore/main/commands/all/api"
	"github.com/xcode75/xcore/main/commands/all/tls"
	"github.com/xcode75/xcore/main/commands/base"
)

// go:generate go run github.com/xcode75/xcore/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
