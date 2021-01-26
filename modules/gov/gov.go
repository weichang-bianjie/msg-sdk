package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/msg-sdk/modules"
)

type govClient struct {
}

func NewClient() Client {
	return govClient{}
}

func (gov govClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgSubmitProposal).Type():
		docMsg := DocTxMsgSubmitProposal{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgVote).Type():
		docMsg := DocTxMsgVote{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgDeposit).Type():
		docMsg := DocTxMsgDeposit{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
