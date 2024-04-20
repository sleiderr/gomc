package game

import (
	"github.com/google/uuid"
)

type GameStatus struct {
	Version      GameVersion     `json:"version"`
	Players      GamePlayers     `json:"players"`
	Description  GameDescription `json:"description"`
	Favicon      string          `json:"favicon,omitempty"`
	SecureChat   bool            `json:"enforcesSecureChat"`
	PreviewsChat bool            `json:"previewsChat"`
}

type GameVersion struct {
	Name     string `json:"name"`
	Protocol uint   `json:"protocol"`
}

type GamePlayers struct {
	Max    uint           `json:"max"`
	Online uint           `json:"online"`
	Sample []OnlinePlayer `json:"sample"`
}

type GameDescription struct {
	Text string `json:"text"`
}

type OnlinePlayer struct {
	Name string
	Id   uuid.UUID
}

func GetGameStatus() *GameStatus {
	version := GameVersion{
		Name:     "1.20.4",
		Protocol: 765,
	}

	players := GamePlayers{
		Max:    16,
		Online: 0,
		Sample: make([]OnlinePlayer, 0),
	}

	description := GameDescription{
		Text: "Welcome to my minecraft server !",
	}

	return &GameStatus{
		version,
		players,
		description,
		"",
		true,
		true,
	}
}
