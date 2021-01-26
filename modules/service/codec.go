package service

import (
	"github.com/irisnet/irismod/modules/service"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
