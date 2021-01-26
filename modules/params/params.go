package params

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/msg-sdk/modules"
)

type paramsClient struct {
}

func NewClient() Client {
	return paramsClient{}
}

func (params paramsClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := false

	return msgDocInfo, ok
}
