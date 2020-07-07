package lol

// LeagueAPI https://developer.riotgames.com/apis#league-v4
type LeagueAPI struct {
	client *Client
}

// APIVersion v4
func (s *LeagueAPI) APIVersion() (version string) {
	return "v4"
}

// APIName league-exp
func (s *LeagueAPI) APIName() (name string) {
	return "league"
}
