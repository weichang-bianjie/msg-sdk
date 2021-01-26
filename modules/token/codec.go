package token

import (
	"github.com/irisnet/irismod/modules/token"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
