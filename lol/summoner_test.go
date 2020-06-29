package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByName(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, _ := client.Summoner.ByName("Kim Chang Ryul")

	assert.Equal("Kim Chang Ryul", res.Name)
}

func TestByPUUID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, _ := client.Summoner.ByPUUID("FJsir2NQU9H6a01tZhHr1TgmwxRg3baLqgMV33tS-MssSnB_mnaYb5EKzJ2XZaMWsE5hBc7tft5oKA")

	assert.Equal("Kim Chang Ryul", res.Name)
}

func TestByAccountID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, _ := client.Summoner.ByAccount("YaoZcfJ4fvVBmWLm8g9fU432yvJSCLBE3nsmDZWhZZQB")

	assert.Equal("Kim Chang Ryul", res.Name)
}

func TestSummonerID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, _ := client.Summoner.SummonerID("aPWJgSeY9bV4Jq6DJ7lOBo3YVr9VvB_fcrdQb3NKllH8WQ")

	assert.Equal("Kim Chang Ryul", res.Name)
}
