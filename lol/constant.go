package lol

//Region
const (
	BR1  = "br1.api.riotgames.com"
	EUN1 = "eun1.api.riotgames.com"
	EUW1 = "euw1.api.riotgames.com"
	JP1  = "jp1.api.riotgames.com"
	KR   = "kr.api.riotgames.com"
	LA1  = "la1.api.riotgames.com"
	LA2  = "la2.api.riotgames.com"
	NA1  = "na1.api.riotgames.com"
	OC1  = "oc1.api.riotgames.com"
	TR1  = "tr1.api.riotgames.com"
	RU   = "ru.api.riotgames.com"
)

// QueueType
type QueueType string

const (
	QueueRankedSolo5X5 = QueueType("RANKED_SOLO_5x5")
	QueueRankedTFT     = QueueType("RANKED_TFT")
	QueueRankedFlexSR  = QueueType("RANKED_FLEX_SR")
	QueueRankedFlexTT  = QueueType("RANKED_FLEX_TT")
)

// TierType
type TierType string

const (
	TierCHALLENGER  = TierType("CHALLENGER")
	TierGRANDMASTER = TierType("GRANDMASTER")
	TierMASTER      = TierType("MASTER")
	TierDIAMOND     = TierType("DIAMOND")
	TierPLATINUM    = TierType("PLATINUM")
	TierGOLD        = TierType("GOLD")
	TierSILVER      = TierType("SILVER")
	TierBRONZE      = TierType("BRONZE")
	TierIRON        = TierType("IRON")
)

type DivisionType string

const (
	DivisionI   = DivisionType("I")
	DivisionII  = DivisionType("II")
	DivisionIII = DivisionType("III")
	DivisionIV  = DivisionType("IV")
)
