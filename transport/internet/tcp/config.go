package tcp

import (
	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
