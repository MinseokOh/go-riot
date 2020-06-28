package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, TestAPIKey, nil)
	res, _ := client.Match.MatchID("4261996769")

	assert.Equal(res.GameID, int64(4261996769))
}
