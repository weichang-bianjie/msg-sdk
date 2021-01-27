package msg_sdk

import (
	"github.com/weichang-bianjie/msg-sdk/codec"
	"github.com/weichang-bianjie/msg-sdk/modules/auth"
	"github.com/weichang-bianjie/msg-sdk/modules/bank"
	"github.com/weichang-bianjie/msg-sdk/modules/coinswap"
	"github.com/weichang-bianjie/msg-sdk/modules/crisis"
	"github.com/weichang-bianjie/msg-sdk/modules/distribution"
	"github.com/weichang-bianjie/msg-sdk/modules/evidence"
	"github.com/weichang-bianjie/msg-sdk/modules/gov"
	"github.com/weichang-bianjie/msg-sdk/modules/htlc"
	"github.com/weichang-bianjie/msg-sdk/modules/ibc"
	"github.com/weichang-bianjie/msg-sdk/modules/nft"
	"github.com/weichang-bianjie/msg-sdk/modules/oracle"
	"github.com/weichang-bianjie/msg-sdk/modules/params"
	"github.com/weichang-bianjie/msg-sdk/modules/random"
	"github.com/weichang-bianjie/msg-sdk/modules/record"
	"github.com/weichang-bianjie/msg-sdk/modules/service"
	"github.com/weichang-bianjie/msg-sdk/modules/slashing"
	"github.com/weichang-bianjie/msg-sdk/modules/staking"
	"github.com/weichang-bianjie/msg-sdk/modules/token"
	"github.com/weichang-bianjie/msg-sdk/modules/upgrade"
)

type MsgClient struct {
	Auth         auth.Client
	Bank         bank.Client
	Staking      staking.Client
	Crisis       crisis.Client
	Distribution distribution.Client
	Evidence     evidence.Client
	Gov          gov.Client
	Ibc          ibc.Client
	Params       params.Client
	Slashing     slashing.Client
	Upgrade      upgrade.Client
	Service      service.Client
	Nft          nft.Client
	Token        token.Client
	Random       random.Client
	Oracle       oracle.Client
	Htlc         htlc.Client
	Record       record.Client
	Coinswap     coinswap.Client
	//Wasm         wasm.Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Auth:         auth.NewClient(),
		Bank:         bank.NewClient(),
		Crisis:       crisis.NewClient(),
		Distribution: distribution.NewClient(),
		Evidence:     distribution.NewClient(),
		Gov:          gov.NewClient(),
		Ibc:          ibc.NewClient(),
		Params:       params.NewClient(),
		Slashing:     slashing.NewClient(),
		Upgrade:      upgrade.NewClient(),
		Staking:      staking.NewClient(),
		Service:      service.NewClient(),
		Nft:          nft.NewClient(),
		Record:       record.NewClient(),
		Random:       random.NewClient(),
		Oracle:       oracle.NewClient(),
		Htlc:         htlc.NewClient(),
		Token:        token.NewClient(),
		Coinswap:     coinswap.NewClient(),
		//Wasm:         wasm.NewClient(),
	}
}
