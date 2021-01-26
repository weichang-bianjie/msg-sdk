package random

import (
	"github.com/irisnet/irismod/modules/random"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(random.AppModuleBasic{})
}
