package staking

import (
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(staking.AppModuleBasic{})
}
