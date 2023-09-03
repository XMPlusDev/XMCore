package core_test

import (
	"testing"

	"github.com/xmplusdev/xmcore/app/dispatcher"
	"github.com/xmplusdev/xmcore/app/proxyman"
	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/net"
	"github.com/xmplusdev/xmcore/common/protocol"
	"github.com/xmplusdev/xmcore/common/serial"
	"github.com/xmplusdev/xmcore/common/uuid"
	. "github.com/xmplusdev/xmcore/core"
	"github.com/xmplusdev/xmcore/features/dns"
	"github.com/xmplusdev/xmcore/features/dns/localdns"
	_ "github.com/xmplusdev/xmcore/main/distro/all"
	"github.com/xmplusdev/xmcore/proxy/dokodemo"
	"github.com/xmplusdev/xmcore/proxy/vmess"
	"github.com/xmplusdev/xmcore/proxy/vmess/outbound"
	"github.com/xmplusdev/xmcore/testing/servers/tcp"
	"google.golang.org/protobuf/proto"
)

func TestXrayDependency(t *testing.T) {
	instance := new(Instance)

	wait := make(chan bool, 1)
	instance.RequireFeatures(func(d dns.Client) {
		if d == nil {
			t.Error("expected dns client fulfilled, but actually nil")
		}
		wait <- true
	})
	instance.AddFeature(localdns.New())
	<-wait
}

func TestXrayClose(t *testing.T) {
	port := tcp.PickPort()

	userID := uuid.New()
	config := &Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
		Inbound: []*InboundHandlerConfig{
			{
				ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
					PortList: &net.PortList{
						Range: []*net.PortRange{net.SinglePortRange(port)},
					},
					Listen: net.NewIPOrDomain(net.LocalHostIP),
				}),
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address: net.NewIPOrDomain(net.LocalHostIP),
					Port:    uint32(0),
					NetworkList: &net.NetworkList{
						Network: []net.Network{net.Network_TCP},
					},
				}),
			},
		},
		Outbound: []*OutboundHandlerConfig{
			{
				ProxySettings: serial.ToTypedMessage(&outbound.Config{
					Receiver: []*protocol.ServerEndpoint{
						{
							Address: net.NewIPOrDomain(net.LocalHostIP),
							Port:    uint32(0),
							User: []*protocol.User{
								{
									Account: serial.ToTypedMessage(&vmess.Account{
										Id: userID.String(),
									}),
								},
							},
						},
					},
				}),
			},
		},
	}

	cfgBytes, err := proto.Marshal(config)
	common.Must(err)

	server, err := StartInstance("protobuf", cfgBytes)
	common.Must(err)
	server.Close()
}
