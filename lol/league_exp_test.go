package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeagueExpEntries(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.LeagueExp.Entries(lol.QueueRankedSolo5X5, lol.TierDIAMOND, lol.DivisionI)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res[0].Tier, string(lol.TierDIAMOND))
}
