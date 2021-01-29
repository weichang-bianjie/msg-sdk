package types

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/types"
)

type (
	Tx    types.Tx
	Block types.Block
	Event abci.Event
)
