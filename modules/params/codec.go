package params

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
