package upgrade

import (
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
