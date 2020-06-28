package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChampionRotations(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, TestAPIKey, nil)
	res, _ := client.Champion.ChampionRotations()

	assert.Equal(res.MaxNewPlayerLevel, 10)
}
