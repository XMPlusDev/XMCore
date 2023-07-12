package core_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/xcode75/xcore/app/dispatcher"
	"github.com/xcode75/xcore/app/proxyman"
	"github.com/xcode75/xcore/common"
	"github.com/xcode75/xcore/common/net"
	"github.com/xcode75/xcore/common/protocol"
	"github.com/xcode75/xcore/common/serial"
	"github.com/xcode75/xcore/common/uuid"
	. "github.com/xcode75/xcore/core"
	"github.com/xcode75/xcore/features/dns"
	"github.com/xcode75/xcore/features/dns/localdns"
	_ "github.com/xcode75/xcore/main/distro/all"
	"github.com/xcode75/xcore/proxy/dokodemo"
	"github.com/xcode75/xcore/proxy/vmess"
	"github.com/xcode75/xcore/proxy/vmess/outbound"
	"github.com/xcode75/xcore/testing/servers/tcp"
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
