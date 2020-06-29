package lol_test

import (
	"fmt"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBySummoner(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.ChampionMastery.BySummoner("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(res[0].ChampionLevel, 7)
}

func TestByChampion(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, _ := client.ChampionMastery.ByChampion("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ", 26)

	assert.Equal(res.ChampionLevel, 7)
}

func TestScore(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, _ := client.ChampionMastery.Scores("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")

	assert.Equal(res, 541)
}
