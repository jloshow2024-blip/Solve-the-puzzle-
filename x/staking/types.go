package staking

import "time"

// StakeRecord represents an individual delegator's stake.
type StakeRecord struct {
	DelegatorAddress string    `json:"delegator_address"`
	StakeAmount      int64     `json:"stake_amount"`
	RewardAmount     int64     `json:"reward_amount"`
	LockUpEndTime    time.Time `json:"lock_up_end_time"`
}

// StakingPool represents global staking pool state.
type StakingPool struct {
	TotalStaked int64   `json:"total_staked"`
	RewardRate  float64 `json:"reward_rate"`
}
