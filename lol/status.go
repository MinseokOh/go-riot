package lol

// SpectatorAPI https://developer.riotgames.com/apis#lol-status-v3/
type StatusAPI struct {
	client *Client
}

// APIVersion v3
func (s *StatusAPI) APIVersion() (version string) {
	return "v3"
}

// APIName status
func (s *StatusAPI) APIName() (name string) {
	return "status"
}

// ShardData Get League of Legends status for the given shard.
func (s *StatusAPI) ShardData() (sharedStatus ShardStatus, err error) {
	path := s.client.GetPath(s, "shard-data")

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &sharedStatus)
	if err != nil {
		return
	}

	return
}
