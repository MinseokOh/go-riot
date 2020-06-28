package lol

//https://developer.riotgames.com/apis#summoner-v4/
type SummonerAPI struct {
	client *Client
}

func (s *SummonerAPI) APIVersion() (version string) {
	return "v4"
}

func (s *SummonerAPI) APIName() (name string) {
	return "summoner"
}

func (s *SummonerAPI) ByAccount(encryptedAccountId string) (summoner Summoner, err error) {
	path := s.client.GetPath(s, "summoners", "by-account", encryptedAccountId)

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

// SummonerID is
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
