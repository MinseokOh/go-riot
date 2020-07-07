package lol

// MatchAPI https://developer.riotgames.com/apis#champion-v3
type MatchAPI struct {
	client *Client
}

// APIVersion v4
func (s *MatchAPI) APIVersion() (version string) {
	return "v4"
}

// APIName match
func (s *MatchAPI) APIName() (name string) {
	return "match"
}

// MatchID Get match by match ID.
func (s *MatchAPI) MatchID(matchID string) (match Match, err error) {
	path := s.client.GetPath(s, "matches", matchID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &match)
	if err != nil {
		return
	}

	return
}

// ByAccountID Get matchlist for games played on given account ID and platform ID and filtered using given filter parameters, if any.
// TODO Optional Parameter
func (s *MatchAPI) ByAccountID(encryptedAccountID string) (match MatchList, err error) {
	path := s.client.GetPath(s, "matchlists", "by-account", encryptedAccountID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &match)
	if err != nil {
		return
	}

	return
}
