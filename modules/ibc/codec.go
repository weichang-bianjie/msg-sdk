package ibc

import (
	ibctransfer "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	"github.com/weichang-bianjie/msg-sdk/codec"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
	)
}
