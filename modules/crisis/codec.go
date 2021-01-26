package crisis

import (
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(crisis.AppModuleBasic{})
}
