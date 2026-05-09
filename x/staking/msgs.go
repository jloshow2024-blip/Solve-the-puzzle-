package staking

// Coin is a minimal coin type used by this scaffold.
type Coin struct {
	Denom  string
	Amount int64
}

// MsgStake is the message to stake tokens.
type MsgStake struct {
	Delegator string
	Amount    Coin
}

// MsgUnstake is the message to unstake tokens.
type MsgUnstake struct {
	Delegator string
	Amount    Coin
}

// NewMsgStake creates a new MsgStake
func NewMsgStake(delegator string, amount Coin) MsgStake {
	return MsgStake{Delegator: delegator, Amount: amount}
}

// NewMsgUnstake creates a new MsgUnstake
func NewMsgUnstake(delegator string, amount Coin) MsgUnstake {
	return MsgUnstake{Delegator: delegator, Amount: amount}
}
