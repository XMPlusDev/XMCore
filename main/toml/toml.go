package toml

import (
	"io"

	"github.com/xmplusdev/xmcore/common"
	"github.com/xmplusdev/xmcore/common/cmdarg"
	"github.com/xmplusdev/xmcore/core"
	"github.com/xmplusdev/xmcore/infra/conf"
	"github.com/xmplusdev/xmcore/infra/conf/serial"
	"github.com/xmplusdev/xmcore/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "TOML",
		Extension: []string{"toml"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					newError("Reading config: ", arg).AtInfo().WriteToLog()
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, newError("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeTOMLConfig(r)
					if err != nil {
						return nil, newError("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadTOMLConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
