package gov

import (
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
		upgrade.AppModuleBasic{},
	)
}
