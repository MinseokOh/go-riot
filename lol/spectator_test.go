package lol_test

import (
	"go-riot"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpectatorFeaturedGames(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Spectator.FeaturedGames()
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.NotEqual(len(res.GameList), 0)
}

func TestSpectatorBySummonerID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)

	featured, err := client.Spectator.FeaturedGames()
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	summoner, err := client.Summoner.ByName(featured.GameList[0].Participants[0].SummonerName)
	if err != nil {
		assert.Fail(err.Error())
	}

	res, err := client.Spectator.BySummonerID(summoner.ID)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.NotEmpty(res)
}
