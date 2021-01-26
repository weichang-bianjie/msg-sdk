package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/weichang-bianjie/msg-sdk/modules"
)

type distributionClient struct {
}

func NewClient() Client {
	return distributionClient{}
}

func (distribution distributionClient) HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.Type() {
	case new(MsgStakeSetWithdrawAddress).Type():
		docMsg := DocTxMsgSetWithdrawAddress{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgWithdrawDelegatorReward).Type():
		docMsg := DocTxMsgWithdrawDelegatorReward{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgWithdrawValidatorCommission).Type():
		docMsg := DocTxMsgWithdrawValidatorCommission{}
		return docMsg.HandleTxMsg(msg), ok
	case new(MsgFundCommunityPool).Type():
		docMsg := DocTxMsgFundCommunityPool{}
		return docMsg.HandleTxMsg(msg), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
