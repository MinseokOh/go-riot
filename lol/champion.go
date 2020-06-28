package lol

//https://developer.riotgames.com/apis#champion-v3
type ChampionAPI struct {
	client *Client
}

func (s *ChampionAPI) APIVersion() (version string) {
	return "v3"
}

func (s *ChampionAPI) APIName() (name string) {
	return "platform"
}

func (s *ChampionAPI) ChampionRotations() (freeChampions FreeChampions, err error) {
	path := s.client.GetPath(s, "champion-rotations")

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &freeChampions)
	if err != nil {
		return
	}

	return
}
