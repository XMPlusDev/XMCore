package command_test

import (
	"context"
	"testing"

	"github.com/xcode75/xcore/app/dispatcher"
	"github.com/xcode75/xcore/app/log"
	. "github.com/xcode75/xcore/app/log/command"
	"github.com/xcode75/xcore/app/proxyman"
	_ "github.com/xcode75/xcore/app/proxyman/inbound"
	_ "github.com/xcode75/xcore/app/proxyman/outbound"
	"github.com/xcode75/xcore/common"
	"github.com/xcode75/xcore/common/serial"
	"github.com/xcode75/xcore/core"
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
