package coinswap

import (
	"github.com/irisnet/irismod/modules/coinswap"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(coinswap.AppModuleBasic{})
}
