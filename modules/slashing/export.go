package slashing

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/msg-sdk/modules"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
