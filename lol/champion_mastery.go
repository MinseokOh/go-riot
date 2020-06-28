package lol

import "strconv"

//https://developer.riotgames.com/apis#champion-v3
type ChampionMasteryAPI struct {
	client *Client
}

func (s *ChampionMasteryAPI) APIVersion() (version string) {
	return "v4"
}

func (s *ChampionMasteryAPI) APIName() (name string) {
	return "champion-mastery"
}

func (s *ChampionMasteryAPI) BySummoner(encryptedSummonerID string) (masteries []ChampionMastery, err error) {
	path := s.client.GetPath(s, "champion-masteries", "by-summoner", encryptedSummonerID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &masteries)
	if err != nil {
		return
	}

	return
}

func (s *ChampionMasteryAPI) ByChampion(encryptedSummonerID string, ChampionID int) (mastery ChampionMastery, err error) {
	path := s.client.GetPath(s, "champion-masteries", "by-summoner", encryptedSummonerID, "by-champion", strconv.Itoa(ChampionID))

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &mastery)
	if err != nil {
		return
	}

	return
}

// - /lol/champion-mastery/v4/scores/by-summoner/{encryptedSummonerId}
func (s *ChampionMasteryAPI) Scores(encryptedSummonerID string) (score int, err error) {
	path := s.client.GetPath(s, "scores", "by-summoner", encryptedSummonerID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &score)
	if err != nil {
		return
	}

	return
}
