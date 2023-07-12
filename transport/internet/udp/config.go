package udp

import (
	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
