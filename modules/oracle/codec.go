package oracle

import (
	"github.com/irisnet/irismod/modules/oracle"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(oracle.AppModuleBasic{})
}
