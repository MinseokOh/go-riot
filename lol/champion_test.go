package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChampionRotations(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Champion.ChampionRotations()
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.MaxNewPlayerLevel, 10)
}
