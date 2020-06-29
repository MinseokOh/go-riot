package lol

type ClashAPI struct {
	client *Client
}

func (s *ClashAPI) APIVersion() (version string) {
	return "v1"
}

func (s *ClashAPI) APIName() (name string) {
	return "clash"
}

func (s *ClashAPI) BySummoner(summonerID string) (players []Player, err error) {
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

func (s *ClashAPI) TournamentsByTeam() (tournaments []Tournament, err error) {
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

func (s *ClashAPI) TournamentsByTournamentID() (tournaments []Tournament, err error) {
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
