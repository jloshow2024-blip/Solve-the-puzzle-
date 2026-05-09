package staking

import "testing"

func TestKeeperStakeUnstake(t *testing.T) {
	k := NewKeeper()
	delegator := "addr1"

	k.Stake(delegator, 100)
	rec, ok := k.GetStakeRecord(delegator)
	if !ok {
		t.Fatalf("expected stake record to exist")
	}
	if rec.StakeAmount != 100 {
		t.Fatalf("expected stake amount 100, got %d", rec.StakeAmount)
	}

	success := k.Unstake(delegator, 40)
	if !success {
		t.Fatalf("unstake failed")
	}

	rec, _ = k.GetStakeRecord(delegator)
	if rec.StakeAmount != 60 {
		t.Fatalf("expected stake amount 60 after unstake, got %d", rec.StakeAmount)
	}

	// Unstake too much
	success = k.Unstake(delegator, 1000)
	if success {
		t.Fatalf("unstake should have failed due to insufficient funds")
	}
}
