package command_test

import (
	"context"
	"testing"

	"github.com/xmplusdev/xmcore/app/dispatcher"
	"github.com/xmplusdev/xmcore/app/log"
	. "github.com/xmplusdev/xmcore/app/log/command"
	"github.com/xmplusdev/xmcore/app/proxyman"
	_ "github.com/xmplusdev/xmcore/app/proxyman/inbound"
	_ "github.com/xmplusdev/xmcore/app/proxyman/outbound"
	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/serial"
	"github.com/xmplusdev/xmcore/core"
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
