package lol

import "strconv"

// ChampionMasteryAPI https://developer.riotgames.com/apis#champion-v3
type ChampionMasteryAPI struct {
	client *Client
}

// APIVersion v4
func (s *ChampionMasteryAPI) APIVersion() (version string) {
	return "v4"
}

// APIName champion-mastery
func (s *ChampionMasteryAPI) APIName() (name string) {
	return "champion-mastery"
}

// BySummoner Get all champion mastery entries sorted by number of champion points descending,
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

// ByChampion Get a champion mastery by player ID and champion ID.
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

// Scores Get a player's total champion mastery score, which is the sum of individual champion mastery levels.
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
