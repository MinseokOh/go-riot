package lol

// SummonerAPI https://developer.riotgames.com/apis#summoner-v4/
type SummonerAPI struct {
	client *Client
}

// APIVersion v4
func (s *SummonerAPI) APIVersion() (version string) {
	return "v4"
}

// APIName summoner
func (s *SummonerAPI) APIName() (name string) {
	return "summoner"
}

// ByAccountID Get a summoner by account ID.
func (s *SummonerAPI) ByAccountID(encryptedAccountID string) (summoner Summoner, err error) {
	path := s.client.GetPath(s, "summoners", "by-account", encryptedAccountID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &summoner)
	if err != nil {
		return
	}

	return
}

// ByName Get a summoner by summoner name.
func (s *SummonerAPI) ByName(name string) (summoner Summoner, err error) {
	path := s.client.GetPath(s, "summoners", "by-name", name)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &summoner)
	if err != nil {
		return
	}

	return
}

// ByPUUID Get a summoner by PUUID.
func (s *SummonerAPI) ByPUUID(encryptedPUUID string) (summoner Summoner, err error) {
	path := s.client.GetPath(s, "summoners", "by-puuid", encryptedPUUID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &summoner)
	if err != nil {
		return
	}

	return
}

// SummonerID Get a summoner by summoner ID.
// *Consistently looking up summoner ids that don't exist will result in a blacklist.*
func (s *SummonerAPI) SummonerID(encryptedSummonerID string) (summoner Summoner, err error) {
	path := s.client.GetPath(s, "summoners", encryptedSummonerID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &summoner)
	if err != nil {
		return
	}

	return
}
