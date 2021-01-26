package nft

import (
	"github.com/irisnet/irismod/modules/nft"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
