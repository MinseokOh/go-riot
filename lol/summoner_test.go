package lol_test

import (
	"go-riot"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummonerByName(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Summoner.ByName("Kim Chang Ryul")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal("Kim Chang Ryul", res.Name)
}

func TestSummonerByPUUID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Summoner.ByPUUID("FJsir2NQU9H6a01tZhHr1TgmwxRg3baLqgMV33tS-MssSnB_mnaYb5EKzJ2XZaMWsE5hBc7tft5oKA")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal("Kim Chang Ryul", res.Name)
}

func TestSummonerByAccountID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Summoner.ByAccountID("YaoZcfJ4fvVBmWLm8g9fU432yvJSCLBE3nsmDZWhZZQB")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal("Kim Chang Ryul", res.Name)
}

func TestSummonerID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Summoner.SummonerID("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal("Kim Chang Ryul", res.Name)
}
