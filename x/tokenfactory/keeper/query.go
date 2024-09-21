package keeper

import (
	"loan/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
