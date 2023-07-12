package outbound

import (
	"context"
	"os"

	"github.com/sagernet/sing/common/uot"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/transport/internet"
	"github.com/xmplusdev/xmcore/transport/internet/stat"
)

func (h *Handler) getUoTConnection(ctx context.Context, dest net.Destination) (stat.Connection, error) {
	if !dest.Address.Family().IsDomain() {
		return nil, os.ErrInvalid
	}
	var uotVersion int
	if dest.Address.Domain() == uot.MagicAddress {
		uotVersion = uot.Version
	} else if dest.Address.Domain() == uot.LegacyMagicAddress {
		uotVersion = uot.LegacyVersion
	} else {
		return nil, os.ErrInvalid
	}
	packetConn, err := internet.ListenSystemPacket(ctx, &net.UDPAddr{IP: net.AnyIP.IP(), Port: 0}, h.streamSettings.SocketSettings)
	if err != nil {
		return nil, newError("unable to listen socket").Base(err)
	}
	conn := uot.NewServerConn(packetConn, uotVersion)
	return h.getStatCouterConnection(conn), nil
}
