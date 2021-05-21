package keeper

import (
	"github.com/ahmetson/random-number-blockchain/x/randomnumberblockchain/types"
)

var _ types.QueryServer = Keeper{}
