package integration

import (
	"fmt"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/weichang-bianjie/msg-sdk/utils"
)

func (s IntegrationTestSuite) TestBank() {
	cases := []SubTest{
		//{
		//	"TestQueryAccount",
		//	queryAccount,
		//},
		{
			"TestSend",
			send,
		},
		//{
		//	"TestMultiSend",
		//	multiSend,
		//},
		//{
		//	"TestSimulate",
		//	simulate,
		//},
		//{
		//	"TestSendWitchSpecAccountInfo",
		//	sendWitchSpecAccountInfo,
		//},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func send(s IntegrationTestSuite) {
	var msg bank.MsgSend
	banksdk := "{\"from_address\":\"iaa174qyl02cupyqq77cqqtdl0frda6dl3rp2h9snu\",\"to_address\":\"iaa16urnygz4rrjxlhy5x8kn3ag7p0ufhukl3utulm\",\"amount\":[{\"denom\":\"ubif\",\"amount\":\"100000000\"}]}"
	utils.UnMarshalJsonIgnoreErr(banksdk, msg)
	if bankDoc, ok := s.Bank.HandleTxMsg(&msg); ok {
		fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
	}
}
