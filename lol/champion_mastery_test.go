package lol_test

import (
	"fmt"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBySummoner(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, TestAPIKey, nil)
	res, err := client.ChampionMastery.BySummoner(TestSummonerEnID)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(res[0].ChampionLevel, 7)
}

func TestByChampion(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, TestAPIKey, nil)
	res, _ := client.ChampionMastery.ByChampion(TestSummonerEnID, 26)

	assert.Equal(res.ChampionLevel, 7)
}

func TestScore(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, TestAPIKey, nil)
	res, _ := client.ChampionMastery.Scores(TestSummonerEnID)

	assert.Equal(res, 541)
}