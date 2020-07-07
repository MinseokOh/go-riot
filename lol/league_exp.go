package lol

// LeagueExpAPI https://developer.riotgames.com/apis#league-exp-v4
type LeagueExpAPI struct {
	client *Client
}

// APIVersion v4
func (s *LeagueExpAPI) APIVersion() (version string) {
	return "v4"
}

// APIName league-exp
func (s *LeagueExpAPI) APIName() (name string) {
	return "league-exp"
}

// Entries Get all the league entries.
// TODO Optional Parm
func (s *LeagueExpAPI) Entries(queue QueueType, tier TierType, division DivisionType) (leagueEntry []LeagueEntry, err error) {
	path := s.client.GetPath(s, "entries", string(queue), string(tier), string(division))

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueEntry)
	if err != nil {
		return
	}

	return
}
