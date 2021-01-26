package slashing

import (
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
}
