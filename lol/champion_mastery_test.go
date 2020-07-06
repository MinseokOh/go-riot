package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasteryBySummoner(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.ChampionMastery.BySummoner("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res[0].ChampionLevel, 7)
}

func TestMasteryByChampion(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.ChampionMastery.ByChampion("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ", 26)
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.ChampionLevel, 7)
}

func TestMasteryScore(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.ChampionMastery.Scores("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res, 541)
}
