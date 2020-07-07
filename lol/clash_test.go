package lol_test

import (
	"go-riot"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClashBySummonerID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Clash.BySummonerID("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	if len(res) == 0 {
		return
	}

	assert.Equal(res[0].AccountID, int64(4261996769))
}

func TestClashByTeamID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Clash.ByTeamID("1761")
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res.ID, "1761")
}

func TestClashTournaments(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Clash.Tournaments()
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res[0].ID, 1761)
}

func TestClashTournamentsByTeamID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Clash.TournamentsByTeamID("1761")
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res.ID, 144)
}

func TestClashTournamentsByTournamentID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Clash.TournamentsByTournamentID("144")
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res.ID, 144)
}
