package lol_test

import (
	"go-riot"
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(riot.KR, "RGAPI-6df8ce4c-c548-44cc-b35f-f06c59f95627", nil)
	res, err := client.Match.MatchID("4261996769")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.GameID, int64(4261996769))
}
