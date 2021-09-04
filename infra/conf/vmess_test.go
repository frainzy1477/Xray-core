package conf_test

import (
	"testing"

	"github.com/frainzy1477/xray-core/common/net"
	"github.com/frainzy1477/xray-core/common/protocol"
	"github.com/frainzy1477/xray-core/common/serial"
	. "github.com/frainzy1477/xray-core/infra/conf"
	"github.com/frainzy1477/xray-core/proxy/vmess"
	"github.com/frainzy1477/xray-core/proxy/vmess/inbound"
	"github.com/frainzy1477/xray-core/proxy/vmess/outbound"
)

func TestVMessOutbound(t *testing.T) {
	creator := func() Buildable {
		return new(VMessOutboundConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"vnext": [{
					"address": "127.0.0.1",
					"port": 80,
					"users": [
						{
							"id": "e641f5ad-9397-41e3-bf1a-e8740dfed019",
							"email": "love@example.com",
							"level": 255
						}
					]
				}]
			}`,
			Parser: loadJSON(creator),
			Output: &outbound.Config{
				Receiver: []*protocol.ServerEndpoint{
					{
						Address: &net.IPOrDomain{
							Address: &net.IPOrDomain_Ip{
								Ip: []byte{127, 0, 0, 1},
							},
						},
						Port: 80,
						User: []*protocol.User{
							{
								Email: "love@example.com",
								Level: 255,
								Account: serial.ToTypedMessage(&vmess.Account{
									Id:      "e641f5ad-9397-41e3-bf1a-e8740dfed019",
									AlterId: 0,
									SecuritySettings: &protocol.SecurityConfig{
										Type: protocol.SecurityType_AUTO,
									},
								}),
							},
						},
					},
				},
			},
		},
	})
}

func TestVMessInbound(t *testing.T) {
	creator := func() Buildable {
		return new(VMessInboundConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"clients": [
					{
						"id": "27848739-7e62-4138-9fd3-098a63964b6b",
						"level": 0,
						"alterId": 16,
						"email": "love@example.com",
						"security": "aes-128-gcm"
					}
				],
				"default": {
					"level": 0,
					"alterId": 32
				},
				"detour": {
					"to": "tag_to_detour"
				},
				"disableInsecureEncryption": true
			}`,
			Parser: loadJSON(creator),
			Output: &inbound.Config{
				User: []*protocol.User{
					{
						Level: 0,
						Email: "love@example.com",
						Account: serial.ToTypedMessage(&vmess.Account{
							Id:      "27848739-7e62-4138-9fd3-098a63964b6b",
							AlterId: 16,
							SecuritySettings: &protocol.SecurityConfig{
								Type: protocol.SecurityType_AES128_GCM,
							},
						}),
					},
				},
				Default: &inbound.DefaultConfig{
					Level:   0,
					AlterId: 32,
				},
				Detour: &inbound.DetourConfig{
					To: "tag_to_detour",
				},
				SecureEncryptionOnly: true,
			},
		},
	})
}
