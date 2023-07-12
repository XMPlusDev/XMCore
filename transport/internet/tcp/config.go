package tcp

import (
	"github.com/xcode75/xcore/common"
	"github.com/xcode75/xcore/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
