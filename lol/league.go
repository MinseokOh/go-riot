package lol

import "go-riot"

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

// ChallengerLeagues Get the challenger league for given queue.
func (s *LeagueAPI) ChallengerLeagues(queue riot.QueueType) (leagueList LeagueList, err error) {
	path := s.client.GetPath(s, "challengerleagues", "by-queue", string(queue))

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueList)
	if err != nil {
		return
	}

	return
}

// GrandMasterLeagues Get the grandmaster league of a specific queue.
func (s *LeagueAPI) GrandMasterLeagues(queue riot.QueueType) (leagueList LeagueList, err error) {
	path := s.client.GetPath(s, "grandmasterleagues", "by-queue", string(queue))

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueList)
	if err != nil {
		return
	}

	return
}

// MasterLeagues Get the master league of a specific queue.
func (s *LeagueAPI) MasterLeagues(queue riot.QueueType) (leagueList LeagueList, err error) {
	path := s.client.GetPath(s, "masterleagues", "by-queue", string(queue))

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueList)
	if err != nil {
		return
	}

	return
}

// EntriesBySummonerID Get league entries in all queues for a given summoner ID.
func (s *LeagueAPI) EntriesBySummonerID(encryptedSummonerID string) (leagueEntries []LeagueEntry, err error) {
	path := s.client.GetPath(s, "entries", "by-summoner", encryptedSummonerID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueEntries)
	if err != nil {
		return
	}

	return
}

// Entries Get all the league entries.
// TODO Optional Param
func (s *LeagueAPI) Entries(queue riot.QueueType, tier riot.TierType, division riot.DivisionType) (leagueEntries []LeagueEntry, err error) {
	path := s.client.GetPath(s, "entries", string(queue), string(tier), string(division))

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueEntries)
	if err != nil {
		return
	}

	return
}

// Leagues Get league with given ID, including inactive entries.
// *Consistently looking up league ids that don't exist will result in a blacklist.*
func (s *LeagueAPI) Leagues(leagueID string) (leagueList LeagueList, err error) {
	path := s.client.GetPath(s, "leagues", leagueID)

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return
	}

	err = s.client.Do(req, &leagueList)
	if err != nil {
		return
	}

	return
}
