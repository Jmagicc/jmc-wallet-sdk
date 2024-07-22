package doge

import (
	"github.com/Jmagicc/jmc-wallet-sdk/core/base"
)

var (
	_ base.Account = (*Account)(nil)
	_ base.Chain   = (*Chain)(nil)
	_ base.Token   = (*Chain)(nil)
	// _ base.Transaction = (*Transaction)(nil)
)
