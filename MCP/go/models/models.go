package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Stadium represents the Stadium schema from the OpenAPI specification
type Stadium struct {
	Country string `json:"Country,omitempty"`
	Geolat float64 `json:"GeoLat,omitempty"`
	Geolong float64 `json:"GeoLong,omitempty"`
	Name string `json:"Name,omitempty"`
	Playingsurface string `json:"PlayingSurface,omitempty"`
	Stadiumid int `json:"StadiumID,omitempty"`
	State string `json:"State,omitempty"`
	City string `json:"City,omitempty"`
	Capacity int `json:"Capacity,omitempty"`
	TypeField string `json:"Type,omitempty"`
}

// Play represents the Play schema from the OpenAPI specification
type Play struct {
	Quartername string `json:"QuarterName,omitempty"`
	Yardline int `json:"YardLine,omitempty"`
	Created string `json:"Created,omitempty"`
	Playid int `json:"PlayID,omitempty"`
	Distance int `json:"Distance,omitempty"`
	Opponent string `json:"Opponent,omitempty"`
	Playstats []PlayStat `json:"PlayStats,omitempty"`
	Quarterid int `json:"QuarterID,omitempty"`
	Sequence int `json:"Sequence,omitempty"`
	Playtime string `json:"PlayTime,omitempty"`
	Yardlineterritory string `json:"YardLineTerritory,omitempty"`
	Isscoringplay bool `json:"IsScoringPlay,omitempty"`
	Timeremainingseconds int `json:"TimeRemainingSeconds,omitempty"`
	Updated string `json:"Updated,omitempty"`
	Description string `json:"Description,omitempty"`
	Yardstoendzone int `json:"YardsToEndZone,omitempty"`
	TypeField string `json:"Type,omitempty"`
	Timeremainingminutes int `json:"TimeRemainingMinutes,omitempty"`
	Yardsgained int `json:"YardsGained,omitempty"`
	Down int `json:"Down,omitempty"`
	Scoringplay ScoringPlay `json:"ScoringPlay,omitempty"`
	Team string `json:"Team,omitempty"`
}

// PlayByPlay represents the PlayByPlay schema from the OpenAPI specification
type PlayByPlay struct {
	Quarters []Quarter `json:"Quarters,omitempty"`
	Score Score `json:"Score,omitempty"`
	Plays []Play `json:"Plays,omitempty"`
}

// PlayStat represents the PlayStat schema from the OpenAPI specification
type PlayStat struct {
	Passingsackyards int `json:"PassingSackYards,omitempty"`
	Opponent string `json:"Opponent,omitempty"`
	Receivingyards int `json:"ReceivingYards,omitempty"`
	Fieldgoalreturnyards int `json:"FieldGoalReturnYards,omitempty"`
	Kickoffs int `json:"Kickoffs,omitempty"`
	Sackyards float64 `json:"SackYards,omitempty"`
	Fumblesforced int `json:"FumblesForced,omitempty"`
	Fieldgoalreturns int `json:"FieldGoalReturns,omitempty"`
	Blockedkickreturnyards int `json:"BlockedKickReturnYards,omitempty"`
	Kickreturnyards int `json:"KickReturnYards,omitempty"`
	Fieldgoalshadblocked int `json:"FieldGoalsHadBlocked,omitempty"`
	Receivingtouchdowns int `json:"ReceivingTouchdowns,omitempty"`
	Solotackles int `json:"SoloTackles,omitempty"`
	Fieldgoalsattempted int `json:"FieldGoalsAttempted,omitempty"`
	Passingyards int `json:"PassingYards,omitempty"`
	Blockedkickreturns int `json:"BlockedKickReturns,omitempty"`
	Receptions int `json:"Receptions,omitempty"`
	Playid int `json:"PlayID,omitempty"`
	Fumblesrecovered int `json:"FumblesRecovered,omitempty"`
	Interceptionreturnyards int `json:"InterceptionReturnYards,omitempty"`
	Fumbles int `json:"Fumbles,omitempty"`
	Kickoffyards int `json:"KickoffYards,omitempty"`
	Puntreturns int `json:"PuntReturns,omitempty"`
	Interceptionreturntouchdowns int `json:"InterceptionReturnTouchdowns,omitempty"`
	Fumblereturntouchdowns int `json:"FumbleReturnTouchdowns,omitempty"`
	Fieldgoalsyards int `json:"FieldGoalsYards,omitempty"`
	Twopointconversionruns int `json:"TwoPointConversionRuns,omitempty"`
	Passingcompletions int `json:"PassingCompletions,omitempty"`
	Twopointconversionreturns int `json:"TwoPointConversionReturns,omitempty"`
	Extrapointsmade int `json:"ExtraPointsMade,omitempty"`
	Passingattempts int `json:"PassingAttempts,omitempty"`
	Puntyards int `json:"PuntYards,omitempty"`
	Homeoraway string `json:"HomeOrAway,omitempty"`
	Interceptions int `json:"Interceptions,omitempty"`
	Created string `json:"Created,omitempty"`
	Passinginterceptions int `json:"PassingInterceptions,omitempty"`
	Team string `json:"Team,omitempty"`
	Rushingtouchdowns int `json:"RushingTouchdowns,omitempty"`
	Extrapointsattempted int `json:"ExtraPointsAttempted,omitempty"`
	Blockedkicks int `json:"BlockedKicks,omitempty"`
	Puntreturntouchdowns int `json:"PuntReturnTouchdowns,omitempty"`
	Fumbleslost int `json:"FumblesLost,omitempty"`
	Rushingattempts int `json:"RushingAttempts,omitempty"`
	Assistedtackles int `json:"AssistedTackles,omitempty"`
	Fieldgoalreturntouchdowns int `json:"FieldGoalReturnTouchdowns,omitempty"`
	Extrapointshadblocked int `json:"ExtraPointsHadBlocked,omitempty"`
	Punttouchbacks int `json:"PuntTouchbacks,omitempty"`
	Sequence int `json:"Sequence,omitempty"`
	Kickreturntouchdowns int `json:"KickReturnTouchdowns,omitempty"`
	Safeties int `json:"Safeties,omitempty"`
	Twopointconversionreceptions int `json:"TwoPointConversionReceptions,omitempty"`
	Twopointconversionpasses int `json:"TwoPointConversionPasses,omitempty"`
	Passingsacks int `json:"PassingSacks,omitempty"`
	Blockedkickreturntouchdowns int `json:"BlockedKickReturnTouchdowns,omitempty"`
	Passesdefended int `json:"PassesDefended,omitempty"`
	Fieldgoalsmade int `json:"FieldGoalsMade,omitempty"`
	Sacks float64 `json:"Sacks,omitempty"`
	Punts int `json:"Punts,omitempty"`
	Twopointconversionattempts int `json:"TwoPointConversionAttempts,omitempty"`
	Penaltyyards int `json:"PenaltyYards,omitempty"`
	Passingtouchdowns int `json:"PassingTouchdowns,omitempty"`
	Kickofftouchbacks int `json:"KickoffTouchbacks,omitempty"`
	Updated string `json:"Updated,omitempty"`
	Name string `json:"Name,omitempty"`
	Playstatid int `json:"PlayStatID,omitempty"`
	Playerid int `json:"PlayerID,omitempty"`
	Receivingtargets int `json:"ReceivingTargets,omitempty"`
	Direction string `json:"Direction,omitempty"`
	Fumblereturnyards int `json:"FumbleReturnYards,omitempty"`
	Kickreturns int `json:"KickReturns,omitempty"`
	Tacklesforloss int `json:"TacklesForLoss,omitempty"`
	Penalties int `json:"Penalties,omitempty"`
	Puntreturnyards int `json:"PuntReturnYards,omitempty"`
	Rushingyards int `json:"RushingYards,omitempty"`
	Puntshadblocked int `json:"PuntsHadBlocked,omitempty"`
}

// Quarter represents the Quarter schema from the OpenAPI specification
type Quarter struct {
	Description string `json:"Description,omitempty"`
	Quarterid int `json:"QuarterID,omitempty"`
	Awayteamscore int `json:"AwayTeamScore,omitempty"`
	Hometeamscore int `json:"HomeTeamScore,omitempty"`
	Updated string `json:"Updated,omitempty"`
	Created string `json:"Created,omitempty"`
	Number int `json:"Number,omitempty"`
	Name string `json:"Name,omitempty"`
	Scoreid int `json:"ScoreID,omitempty"`
}

// Score represents the Score schema from the OpenAPI specification
type Score struct {
	Awayscore int `json:"AwayScore,omitempty"`
	Awayscorequarter4 int `json:"AwayScoreQuarter4,omitempty"`
	Awayteammoneyline int `json:"AwayTeamMoneyLine,omitempty"`
	Closed bool `json:"Closed,omitempty"`
	Distance string `json:"Distance,omitempty"`
	Globalhometeamid int `json:"GlobalHomeTeamID,omitempty"`
	Homescorequarter1 int `json:"HomeScoreQuarter1,omitempty"`
	Awayscorequarter3 int `json:"AwayScoreQuarter3,omitempty"`
	Awayteamid int `json:"AwayTeamID,omitempty"`
	Overpayout int `json:"OverPayout,omitempty"`
	Channel string `json:"Channel,omitempty"`
	Yardline int `json:"YardLine,omitempty"`
	Isovertime bool `json:"IsOvertime,omitempty"`
	Season int `json:"Season,omitempty"`
	Globalgameid int `json:"GlobalGameID,omitempty"`
	Forecasttemphigh int `json:"ForecastTempHigh,omitempty"`
	Neutralvenue bool `json:"NeutralVenue,omitempty"`
	Forecastwindspeed int `json:"ForecastWindSpeed,omitempty"`
	Homescoreovertime int `json:"HomeScoreOvertime,omitempty"`
	Awayrotationnumber int `json:"AwayRotationNumber,omitempty"`
	Stadiumid int `json:"StadiumID,omitempty"`
	Possession string `json:"Possession,omitempty"`
	Seasontype int `json:"SeasonType,omitempty"`
	Homescorequarter3 int `json:"HomeScoreQuarter3,omitempty"`
	Lastplay string `json:"LastPlay,omitempty"`
	Forecasttemplow int `json:"ForecastTempLow,omitempty"`
	Datetimeutc string `json:"DateTimeUTC,omitempty"`
	Has2ndquarterstarted bool `json:"Has2ndQuarterStarted,omitempty"`
	Date string `json:"Date,omitempty"`
	Datetime string `json:"DateTime,omitempty"`
	Pointspreadhometeammoneyline int `json:"PointSpreadHomeTeamMoneyLine,omitempty"`
	Down int `json:"Down,omitempty"`
	Geolong float64 `json:"GeoLong,omitempty"`
	Canceled bool `json:"Canceled,omitempty"`
	Status string `json:"Status,omitempty"`
	Pointspread float64 `json:"PointSpread,omitempty"`
	Hometeammoneyline int `json:"HomeTeamMoneyLine,omitempty"`
	Globalawayteamid int `json:"GlobalAwayTeamID,omitempty"`
	Homescorequarter2 int `json:"HomeScoreQuarter2,omitempty"`
	Awaytimeouts int `json:"AwayTimeouts,omitempty"`
	Forecastdescription string `json:"ForecastDescription,omitempty"`
	Quarter string `json:"Quarter,omitempty"`
	Hometeamid int `json:"HomeTeamID,omitempty"`
	Isinprogress bool `json:"IsInProgress,omitempty"`
	Awayscorequarter1 int `json:"AwayScoreQuarter1,omitempty"`
	Has1stquarterstarted bool `json:"Has1stQuarterStarted,omitempty"`
	Lastupdated string `json:"LastUpdated,omitempty"`
	Geolat float64 `json:"GeoLat,omitempty"`
	Awayteam string `json:"AwayTeam,omitempty"`
	Awayscoreovertime int `json:"AwayScoreOvertime,omitempty"`
	Refereeid int `json:"RefereeID,omitempty"`
	Overunder float64 `json:"OverUnder,omitempty"`
	Homescore int `json:"HomeScore,omitempty"`
	Homerotationnumber int `json:"HomeRotationNumber,omitempty"`
	Week int `json:"Week,omitempty"`
	Yardlineterritory string `json:"YardLineTerritory,omitempty"`
	Scoreid int `json:"ScoreID,omitempty"`
	Underpayout int `json:"UnderPayout,omitempty"`
	Redzone string `json:"RedZone,omitempty"`
	Pointspreadawayteammoneyline int `json:"PointSpreadAwayTeamMoneyLine,omitempty"`
	Attendance int `json:"Attendance,omitempty"`
	Downanddistance string `json:"DownAndDistance,omitempty"`
	Has3rdquarterstarted bool `json:"Has3rdQuarterStarted,omitempty"`
	Awayscorequarter2 int `json:"AwayScoreQuarter2,omitempty"`
	Quarterdescription string `json:"QuarterDescription,omitempty"`
	Day string `json:"Day,omitempty"`
	Isover bool `json:"IsOver,omitempty"`
	Timeremaining string `json:"TimeRemaining,omitempty"`
	Homescorequarter4 int `json:"HomeScoreQuarter4,omitempty"`
	Gamekey string `json:"GameKey,omitempty"`
	Hasstarted bool `json:"HasStarted,omitempty"`
	Hometimeouts int `json:"HomeTimeouts,omitempty"`
	Stadiumdetails Stadium `json:"StadiumDetails,omitempty"`
	Has4thquarterstarted bool `json:"Has4thQuarterStarted,omitempty"`
	Hometeam string `json:"HomeTeam,omitempty"`
	Gameenddatetime string `json:"GameEndDateTime,omitempty"`
	Forecastwindchill int `json:"ForecastWindChill,omitempty"`
}

// ScoringPlay represents the ScoringPlay schema from the OpenAPI specification
type ScoringPlay struct {
	Seasontype int `json:"SeasonType,omitempty"`
	Gamekey string `json:"GameKey,omitempty"`
	Season int `json:"Season,omitempty"`
	Homescore int `json:"HomeScore,omitempty"`
	Sequence int `json:"Sequence,omitempty"`
	Timeremaining string `json:"TimeRemaining,omitempty"`
	Scoreid int `json:"ScoreID,omitempty"`
	Week int `json:"Week,omitempty"`
	Awayscore int `json:"AwayScore,omitempty"`
	Hometeam string `json:"HomeTeam,omitempty"`
	Awayteam string `json:"AwayTeam,omitempty"`
	Scoringplayid int `json:"ScoringPlayID,omitempty"`
	Date string `json:"Date,omitempty"`
	Quarter string `json:"Quarter,omitempty"`
	Playdescription string `json:"PlayDescription,omitempty"`
	Team string `json:"Team,omitempty"`
}
