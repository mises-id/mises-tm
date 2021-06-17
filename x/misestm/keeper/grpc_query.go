package keeper

import (
	"github.com/mises-id/mises-tm/x/misestm/types"
)

var _ types.QueryServer = Keeper{}
