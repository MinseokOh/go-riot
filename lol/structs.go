package lol

type Response struct {
	Status Status `json:"status"`
}

type Status struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// SUMMONER API

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

// CHAMPION API

type FreeChampions struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

// CHAMPION MASTERY API

type ChampionMastery struct {
	ChampionID                   int    `json:"championId"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	TokensEarned                 int    `json:"tokensEarned"`
	SummonerID                   string `json:"summonerId"`
}

// MATCH API

type TeamBan struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type TeamStats struct {
	TeamID               int       `json:"teamId"`
	Win                  string    `json:"win"`
	FirstBlood           bool      `json:"firstBlood"`
	FirstTower           bool      `json:"firstTower"`
	FirstInhibitor       bool      `json:"firstInhibitor"`
	FirstBaron           bool      `json:"firstBaron"`
	FirstDragon          bool      `json:"firstDragon"`
	FirstRiftHerald      bool      `json:"firstRiftHerald"`
	TowerKills           int       `json:"towerKills"`
	InhibitorKills       int       `json:"inhibitorKills"`
	BaronKills           int       `json:"baronKills"`
	DragonKills          int       `json:"dragonKills"`
	VilemawKills         int       `json:"vilemawKills"`
	RiftHeraldKills      int       `json:"riftHeraldKills"`
	DominionVictoryScore int       `json:"dominionVictoryScore"`
	Bans                 []TeamBan `json:"bans"`
}

type Rune struct {
	RuneID int `json:"runeId"`
	Rank   int `json:"rank"`
}

type Mastery struct {
	MasteryID int `json:masteryId`
	Rank      int `json:"rank"`
}

type Participant struct {
	ParticipantID             int                 `json:"participantId"`
	Runes                     []Rune              `json:"runes"`
	Stats                     ParticipantStats    `json:stats`
	TeamID                    int                 `json:"teamId"`
	ChampionID                int                 `json:"championId"`
	Spell1ID                  int                 `json:"spell1Id"`
	Spell2ID                  int                 `json:"spell2Id"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
	TimeLine                  ParticipantTimeLine `json:"timeline"`
}

type ParticipantStats struct {
	ParticipantID                  int  `json:"participantId"`
	Win                            bool `json:"win"`
	Item0                          int  `json:"item0"`
	Item1                          int  `json:"item1"`
	Item2                          int  `json:"item2"`
	Item3                          int  `json:"item3"`
	Item4                          int  `json:"item4"`
	Item5                          int  `json:"item5"`
	Item6                          int  `json:"item6"`
	Kills                          int  `json:"kills"`
	Deaths                         int  `json:"deaths"`
	Assists                        int  `json:"assists"`
	LargestKillingSpree            int  `json:"largestKillingSpree"`
	LargestMultiKill               int  `json:"largestMultiKill"`
	KillingSprees                  int  `json:"killingSprees"`
	LongestTimeSpentLiving         int  `json:"longestTimeSpentLiving"`
	DoubleKills                    int  `json:"doubleKills"`
	TripleKills                    int  `json:"tripleKills"`
	QuadraKills                    int  `json:"quadraKills"`
	PentaKills                     int  `json:"pentaKills"`
	UnrealKills                    int  `json:"unrealKills"`
	TotalDamageDealt               int  `json:"totalDamageDealt"`
	MagicDamageDealt               int  `json:"magicDamageDealt"`
	PhysicalDamageDealt            int  `json:"physicalDamageDealt"`
	TrueDamageDealt                int  `json:"trueDamageDealt"`
	LargestCriticalStrike          int  `json:"largestCriticalStrike"`
	TotalDamageDealtToChampions    int  `json:"totalDamageDealtToChampions"`
	MagicDamageDealtToChampions    int  `json:"magicDamageDealtToChampions"`
	PhysicalDamageDealtToChampions int  `json:"physicalDamageDealtToChampions"`
	TrueDamageDealtToChampions     int  `json:"trueDamageDealtToChampions"`
	TotalHeal                      int  `json:"totalHeal"`
	TotalUnitsHealed               int  `json:"totalUnitsHealed"`
	DamageSelfMitigated            int  `json:"damageSelfMitigated"`
	DamageDealtToObjectives        int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets           int  `json:"damageDealtToTurrets"`
	VisionScore                    int  `json:"visionScore"`
	TimeCCingOthers                int  `json:"timeCCingOthers"`
	TotalDamageTaken               int  `json:"totalDamageTaken"`
	MagicalDamageTaken             int  `json:"magicalDamageTaken"`
	PhysicalDamageTaken            int  `json:"physicalDamageTaken"`
	TrueDamageTaken                int  `json:"trueDamageTaken"`
	GoldEarned                     int  `json:"goldEarned"`
	GoldSpent                      int  `json:"goldSpent"`
	TurretKills                    int  `json:"turretKills"`
	InhibitorKills                 int  `json:"inhibitorKills"`
	TotalMinionsKilled             int  `json:"totalMinionsKilled"`
	NeutralMinionsKilled           int  `json:"neutralMinionsKilled"`
	TotalTimeCrowdControlDealt     int  `json:"totalTimeCrowdControlDealt"`
	ChampLevel                     int  `json:"champLevel"`
	VisionWardsBoughtInGame        int  `json:"visionWardsBoughtInGame"`
	SightWardsBoughtInGame         int  `json:"sightWardsBoughtInGame"`
	FirstBloodKill                 bool `json:"firstBloodKill"`
	FirstBloodAssist               bool `json:"firstBloodAssist"`
	FirstTowerKill                 bool `json:"firstTowerKill"`
	FirstTowerAssist               bool `json:"firstTowerAssist"`
	FirstInhibitorKill             bool `json:"firstInhibitorKill"`
	FirstInhibitorAssist           bool `json:"firstInhibitorAssist"`
	CombatPlayerScore              int  `json:"combatPlayerScore"`
	ObjectivePlayerScore           int  `json:"objectivePlayerScore"`
	TotalPlayerScore               int  `json:"totalPlayerScore"`
	TotalScoreRank                 int  `json:"totalScoreRank"`
	PlayerScore0                   int  `json:"playerScore0"`
	PlayerScore1                   int  `json:"playerScore1"`
	PlayerScore2                   int  `json:"playerScore2"`
	PlayerScore3                   int  `json:"playerScore3"`
	PlayerScore4                   int  `json:"playerScore4"`
	PlayerScore5                   int  `json:"playerScore5"`
	PlayerScore6                   int  `json:"playerScore6"`
	PlayerScore7                   int  `json:"playerScore7"`
	PlayerScore8                   int  `json:"playerScore8"`
	PlayerScore9                   int  `json:"playerScore9"`
	Perk0                          int  `json:"perk0"`
	Perk0Var1                      int  `json:"perk0Var1"`
	Perk0Var2                      int  `json:"perk0Var2"`
	Perk0Var3                      int  `json:"perk0Var3"`
	Perk1                          int  `json:"perk1"`
	Perk1Var1                      int  `json:"perk1Var1"`
	Perk1Var2                      int  `json:"perk1Var2"`
	Perk1Var3                      int  `json:"perk1Var3"`
	Perk2                          int  `json:"perk2"`
	Perk2Var1                      int  `json:"perk2Var1"`
	Perk2Var2                      int  `json:"perk2Var2"`
	Perk2Var3                      int  `json:"perk2Var3"`
	Perk3                          int  `json:"perk3"`
	Perk3Var1                      int  `json:"perk3Var1"`
	Perk3Var2                      int  `json:"perk3Var2"`
	Perk3Var3                      int  `json:"perk3Var3"`
	Perk4                          int  `json:"perk4"`
	Perk4Var1                      int  `json:"perk4Var1"`
	Perk4Var2                      int  `json:"perk4Var2"`
	Perk4Var3                      int  `json:"perk4Var3"`
	Perk5                          int  `json:"perk5"`
	Perk5Var1                      int  `json:"perk5Var1"`
	Perk5Var2                      int  `json:"perk5Var2"`
	Perk5Var3                      int  `json:"perk5Var3"`
	PerkPrimaryStyle               int  `json:"perkPrimaryStyle"`
	PerkSubStyle                   int  `json:"perkSubStyle"`
	StatPerk0                      int  `json:"statPerk0"`
	StatPerk1                      int  `json:"statPerk1"`
	StatPerk2                      int  `json:"statPerk2"`
}

type ParticipantTimeLine struct {
	ParticipantID int    `json:"participantId"`
	Role          string `json:"role"`
	Lane          string `json:"lane"`

	CsDiffPerMinDeltas          map[string]float32 `json:"csDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float32 `json:"damageTakenPerMinDeltas"`
	DamageTakenDiffPerMinDeltas map[string]float32 `json:"damageTakenDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float32 `json:"creepsPerMinDeltas"`
	XpPerMinDeltas              map[string]float32 `json:"xpPerMinDeltas"`
	XpDiffPerMinDeltas          map[string]float32 `json:"xpDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float32 `json:"goldPerMinDeltas"`
}

type ParticipantIdentity struct {
	ParticipantID int      `json:"participantId"`
	Player        []Player `json:"player"`
}

type Player struct {
	PlatformID        string `json:"platformId"`
	AccountID         string `json:"accountId"`
	SummonerName      string `json:"summonerName"`
	SummonerID        string `json:"summonerId"`
	CurrentPlatformID string `json:"currentPlatformId"`
	CurrentAccountID  string `json:"currentAccountId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	ProfileIcon       int    `json:"profileIcon"`
	Role              string `json:"role"`
	Position          string `json:"position"`
	teamID            string `json:"teamId"`
}

type Match struct {
	GameID                int64                 `json:"gameId"`
	PlatformID            string                `json:"platformId"`
	GameCreation          int64                 `json:"gameCreation"`
	GameDuration          int64                 `json:"gameDuration"`
	QueueID               int                   `json:"queueId"`
	MapID                 int                   `json:"mapId"`
	SeasonID              int                   `json:"seasonId"`
	GameVersion           string                `json:"gameVersion"`
	GameMode              string                `json:"gameMode"`
	GameType              string                `json:"gameType"`
	Teams                 []TeamStats           `json:"teams"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json"participants"`
}

type MatchList struct {
	StartIndex int              `json:"startIndex"`
	TotalGames int              `json:"totalGames"`
	EndIndex   int              `json:"endIndex"`
	Matches    []MatchReference `json:"matches"`
}

type MatchReference struct {
	PlatformID string `json:"platformId"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	Queue      int    `json:"queue"`
	Season     int    `json:"season"`
	Timestamp  int64  `json:"timestamp"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
}

//CLASH API

type Team struct {
	ID           string   `json:"id"`
	TournamentID int      `json:"tournamentId"`
	Name         string   `json:"name"`
	IconID       int      `json:"iconId"`
	Tier         int      `json:"tier"`
	Captain      string   `json:"captain"`
	Abbreviation string   `json:"abbreviation"`
	Players      []Player `json:"players"`
}

type Tournament struct {
	ID               int               `json:"id"`
	ThemeId          int               `json:"themeId"`
	NameKey          string            `json:"nameKey"`
	NameKeySecondary string            `json:"nameKeySecondary"`
	Schedule         []TournamentPhase `json:"schedule"`
}

type TournamentPhase struct {
	ID               int     `json:"id"`
	RegistrationTime float32 `json:"registrationTime"`
	StartTime        float32 `json:"startTime"`
	Cancelled        bool    `json:"cancelled"`
}
