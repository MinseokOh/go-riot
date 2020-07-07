## go-Riot

[![Go Report Card](https://goreportcard.com/badge/github.com/MinseokOh/go-riot)](https://goreportcard.com/report/github.com/MinseokOh/go-riot) [![Build Status](https://travis-ci.com/MinseokOh/go-riot.svg?branch=master)](https://travis-ci.com/MinseokOh/go-riot)

## Installation

```
$ go get github.com/MinseokOh/go-Riot
```

## League Of Legends

### API Coverage

API | Version | Coverage
--- | ------- | --------
Champion | v3 | 100%
Champion-Mastery | v4 | 100%
Clash | v1 | 100%
League | v4 | 100%
League-EXP | v4 | 100%
LOL-Status | v3 | 100%
Match | v4 | 70%
Spectator | v4 | 100%
Summoner | v4 | 100%

### Usage

```
import (
	"fmt"

	"github.com/MinseokOh/go-Riot/lol"
)

func GetSummoner() {
	client := lol.NewClient(lol.KR, APISecret, nil)

	summoner, err := client.Summoner.ByName("DWG ShowMaker")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(summoner.AccountID, summoner.ID, summoner.Name, summoner.PUUID, summoner.PUUID, summoner.SummonerLevel)
}

```

## Teamfight Tactics

## Legends Of Runtera