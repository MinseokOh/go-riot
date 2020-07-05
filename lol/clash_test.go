package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClashBySummoner(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.Clash.BySummoner("4261996769")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res[0].AccountID, int64(4261996769))
}

func TestClashByTeamID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.Clash.ByTeamID("4261996769")
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res.ID, int64(4261996769))
}

func TestClashTournaments(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.Clash.Tournaments()
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res[0].ID, int64(4261996769))
}

func TestClashTournamentsByTeam(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.Clash.TournamentsByTeam("")
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res[0].ID, int64(4261996769))
}

func TestClashTournamentsByTournamentID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.Clash.TournamentsByTournamentID("")
	if err != nil {
		assert.Fail(err.Error())
		return
	}
	assert.Equal(res[0].ID, int64(4261996769))
}
