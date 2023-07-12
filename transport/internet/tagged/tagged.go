package tagged

import (
	"context"

	"github.com/xcode75/xcore/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
