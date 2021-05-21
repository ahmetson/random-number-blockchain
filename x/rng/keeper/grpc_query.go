package keeper

import (
	"github.com/ahmetson/rng/x/rng/types"
)

var _ types.QueryServer = Keeper{}
