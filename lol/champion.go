package lol

// ChampionAPI https://developer.riotgames.com/apis#champion-v3
type ChampionAPI struct {
	client *Client
}

// APIVersion v3
func (s *ChampionAPI) APIVersion() (version string) {
	return "v3"
}

// APIName platform
func (s *ChampionAPI) APIName() (name string) {
	return "platform"
}

// ChampionRotations Returns champion rotations, including free-to-play and low-level free-to-play rotations (REST)
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
