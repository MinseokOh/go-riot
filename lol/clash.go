package lol

// ClashAPI https://developer.riotgames.com/apis#clash-v1
type ClashAPI struct {
	client *Client
}

// APIVersion v1
func (s *ClashAPI) APIVersion() (version string) {
	return "v1"
}

// APIName clash
func (s *ClashAPI) APIName() (name string) {
	return "clash"
}

// BySummonerID Get players by summoner ID.
func (s *ClashAPI) BySummonerID(summonerID string) (players []Player, err error) {
	path := s.client.GetPath(s, "players", "by-summoner", summonerID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &players)
	if err != nil {
		return
	}

	return
}

// ByTeamID Get team by ID.
func (s *ClashAPI) ByTeamID(teamID string) (team Team, err error) {
	path := s.client.GetPath(s, "teams", teamID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &team)
	if err != nil {
		return
	}

	return
}

// Tournaments Get all active or upcoming tournaments.
func (s *ClashAPI) Tournaments() (tournaments []Tournament, err error) {
	path := s.client.GetPath(s, "tournaments")

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &tournaments)
	if err != nil {
		return
	}

	return
}

// TournamentsByTeamID Get tournament by team ID.
func (s *ClashAPI) TournamentsByTeamID(teamID string) (tournament Tournament, err error) {
	path := s.client.GetPath(s, "tournaments", "by-team", teamID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &tournament)
	if err != nil {
		return
	}

	return
}

// TournamentsByTournamentID Get tournament by ID.
func (s *ClashAPI) TournamentsByTournamentID(tournamentId string) (tournament Tournament, err error) {
	path := s.client.GetPath(s, "tournaments", tournamentId)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &tournament)
	if err != nil {
		return
	}

	return
}
