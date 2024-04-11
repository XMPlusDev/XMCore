package all

import (
	"github.com/xmplusdev/xmcore/main/commands/all/api"
	"github.com/xmplusdev/xmcore/main/commands/all/tls"
	"github.com/xmplusdev/xmcore/main/commands/base"
)

// go:generate go run github.com/xmplusdev/xmcore/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
