package distribution

import (
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(distribution.AppModuleBasic{})
}
