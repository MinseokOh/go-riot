package lol

// SpectatorAPI https://developer.riotgames.com/apis#spectator-v4
type SpectatorAPI struct {
	client *Client
}

// APIVersion v4
func (s *SpectatorAPI) APIVersion() (version string) {
	return "v4"
}

// APIName match
func (s *SpectatorAPI) APIName() (name string) {
	return "spectator"
}

// BySummonerID Get current game information for the given summoner ID.
func (s *SpectatorAPI) BySummonerID(encrypedSummonerID string) (currentGameInfo CurrentGameInfo, err error) {
	path := s.client.GetPath(s, "active-games", "by-summoner", encrypedSummonerID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &currentGameInfo)
	if err != nil {
		return
	}

	return
}

// FeaturedGames Get list of featured games.
func (s *SpectatorAPI) FeaturedGames() (featuredGames FeaturedGames, err error) {
	path := s.client.GetPath(s, "featured-games")

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &featuredGames)
	if err != nil {
		return
	}

	return
}
