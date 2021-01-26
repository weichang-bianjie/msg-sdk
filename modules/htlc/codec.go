package htlc

import (
	"github.com/irisnet/irismod/modules/htlc"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(htlc.AppModuleBasic{})
}
