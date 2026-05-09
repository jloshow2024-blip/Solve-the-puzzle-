package staking

import "sync"

// Keeper is a simple in-memory keeper scaffold for staking state.
type Keeper struct {
	mtx        sync.RWMutex
	stakes     map[string]StakeRecord
	stakingPool StakingPool
}

// NewKeeper creates a new Keeper instance.
func NewKeeper() *Keeper {
	return &Keeper{
		stakes: make(map[string]StakeRecord),
		stakingPool: StakingPool{
			TotalStaked: 0,
			RewardRate:  0.05, // default 5% annual for scaffold
		},
	}
}

// Stake adds amount to a delegator's stake.
func (k *Keeper) Stake(delegator string, amount int64) {
	k.mtx.Lock()
	defer k.mtx.Unlock()
	rec, ok := k.stakes[delegator]
	if !ok {
		rec = StakeRecord{DelegatorAddress: delegator}
	}
	rec.StakeAmount += amount
	k.stakes[delegator] = rec
	k.stakingPool.TotalStaked += amount
}

// Unstake removes amount from a delegator's stake if available.
func (k *Keeper) Unstake(delegator string, amount int64) bool {
	k.mtx.Lock()
	defer k.mtx.Unlock()
	rec, ok := k.stakes[delegator]
	if !ok || rec.StakeAmount < amount {
		return false
	}
	rec.StakeAmount -= amount
	k.stakes[delegator] = rec
	k.stakingPool.TotalStaked -= amount
	return true
}

// GetStakeRecord returns a copy of the stake record for a delegator.
func (k *Keeper) GetStakeRecord(delegator string) (StakeRecord, bool) {
	k.mtx.RLock()
	defer k.mtx.RUnlock()
	rec, ok := k.stakes[delegator]
	return rec, ok
}
