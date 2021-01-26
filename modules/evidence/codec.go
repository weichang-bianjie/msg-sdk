package evidence

import (
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(evidence.AppModuleBasic{})
}
