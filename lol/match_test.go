package lol_test

import (
	"go-riot/lol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchID(t *testing.T) {
	assert := assert.New(t)

	client := lol.NewClient(lol.KR, "RGAPI-c6b73352-2c6b-445a-b023-32c46a9d939d", nil)
	res, err := client.Match.MatchID("4261996769")
	if err != nil {
		assert.Fail(err.Error())
		return
	}

	assert.Equal(res.GameID, int64(4261996769))
}
