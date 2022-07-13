package model

import (
	"math/rand"
	"time"

	constants "github.com/nvzard/soccer-manager/helpers"
)

type Player struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Country     string    `json:"country"`
	Age         uint8     `json:"age"`
	Position    string    `json:"position"`
	MarketValue int       `json:"marketValue" gorm:"default:1000000"`
	TeamID      uint      `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type PlayerPatch struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
}

func (player *Player) GeneratePlayer(position string) {
	player.FirstName = constants.GetRandomFirstName()
	player.LastName = constants.GetRandomLastName()
	player.Country = constants.GetRandomCountry()
	player.Age = constants.GetRandomAge()
	player.Position = position
}

func (player *Player) TransferPlayer(teamID uint) {
	player.TeamID = teamID
	player.increaseMarketValue()
}

// Randomly increase player market value b/w 10 and 100 percent
func (player *Player) increaseMarketValue() {
	rand.Seed(time.Now().UnixNano())
	increasePercentage := 10 + rand.Intn(100-10)
	player.MarketValue += player.MarketValue * (increasePercentage / 100)
}

func GeneratePlayers(goalkeeperCount, defendersCount, midfieldersCount, attackersCount int) []Player {
	totalPlayers := goalkeeperCount + defendersCount + midfieldersCount + attackersCount
	players := make([]Player, totalPlayers)
	i := 0

	for goalkeeperCount > 0 {
		players[i].GeneratePlayer(constants.GOALKEEPER)
		goalkeeperCount--
		i++
	}

	for defendersCount > 0 {
		players[i].GeneratePlayer(constants.DEFENDER)
		defendersCount--
		i++
	}

	for midfieldersCount > 0 {
		players[i].GeneratePlayer(constants.MIDFIELDER)
		midfieldersCount--
		i++
	}

	for attackersCount > 0 {
		players[i].GeneratePlayer(constants.ATTACKER)
		attackersCount--
		i++
	}

	return players
}
