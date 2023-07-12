package udp

import (
	"github.com/xcode75/xcore/common"
	"github.com/xcode75/xcore/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
