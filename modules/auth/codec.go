package auth

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
