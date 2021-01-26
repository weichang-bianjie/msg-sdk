package bank

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(bank.AppModuleBasic{})
}
