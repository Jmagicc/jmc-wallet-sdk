package aptos

import (
	"github.com/Jmagicc/jmc-wallet-sdk/core/base"
)

var (
	_ base.Account     = (*Account)(nil)
	_ base.Chain       = (*Chain)(nil)
	_ base.Token       = (*Token)(nil)
	_ base.Transaction = (*Transaction)(nil)

	_ base.SignedTransaction = (*SignedTransaction)(nil)
)
