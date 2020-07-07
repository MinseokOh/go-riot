package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeagueChallengerLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.ChallengerLeagues(lol.QueueRankedSolo5X5)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.Tier, string(lol.TierCHALLENGER))
}

func TestLeagueGrandMasterLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.GrandMasterLeagues(lol.QueueRankedSolo5X5)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.Tier, string(lol.TierGRANDMASTER))
}
func TestLeagueMasterLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.MasterLeagues(lol.QueueRankedSolo5X5)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.Tier, string(lol.TierMASTER))
}

func TestLeagueEntries(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.Entries(lol.QueueRankedSolo5X5, lol.TierDIAMOND, lol.DivisionI)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res[0].Tier, string(lol.TierDIAMOND))
}

func TestLeagueEntriesBySummonerID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	summoner, err := client.Summoner.ByName("DWG ShowMaker")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	res, err := client.League.EntriesBySummonerID(summoner.ID)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res[0].Tier, string(lol.TierCHALLENGER))
}

func TestLeagueLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	summoner, err := client.Summoner.ByName("DWG ShowMaker")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	entries, err := client.League.EntriesBySummonerID(summoner.ID)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	res, err := client.League.Leagues(entries[0].LeagueID)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.LeagueID, entries[0].LeagueID)
}
