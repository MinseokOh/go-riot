package lol_test

import (
	"go-riot"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeagueChallengerLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.ChallengerLeagues(riot.QueueRankedSolo5X5)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.Tier, string(riot.TierCHALLENGER))
}

func TestLeagueGrandMasterLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.GrandMasterLeagues(riot.QueueRankedSolo5X5)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.Tier, string(riot.TierGRANDMASTER))
}
func TestLeagueMasterLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.MasterLeagues(riot.QueueRankedSolo5X5)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.Tier, string(riot.TierMASTER))
}

func TestLeagueEntries(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.League.Entries(riot.QueueRankedSolo5X5, riot.TierDIAMOND, riot.DivisionI)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res[0].Tier, string(riot.TierDIAMOND))
}

func TestLeagueEntriesBySummonerID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
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

	assert.Equal(res[0].Tier, string(riot.TierCHALLENGER))
}

func TestLeagueLeagues(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
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
