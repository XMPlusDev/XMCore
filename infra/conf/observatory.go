package conf

import (
	"github.com/golang/protobuf/proto"
	"github.com/xcode75/xcore/app/observatory"
	"github.com/xcode75/xcore/infra/conf/cfgcommon/duration"
)

type ObservatoryConfig struct {
	SubjectSelector   []string          `json:"subjectSelector"`
	ProbeURL          string            `json:"probeURL"`
	ProbeInterval     duration.Duration `json:"probeInterval"`
	EnableConcurrency bool              `json:"enableConcurrency"`
}

func (o *ObservatoryConfig) Build() (proto.Message, error) {
	return &observatory.Config{SubjectSelector: o.SubjectSelector, ProbeUrl: o.ProbeURL, ProbeInterval: int64(o.ProbeInterval), EnableConcurrency: o.EnableConcurrency}, nil
}
