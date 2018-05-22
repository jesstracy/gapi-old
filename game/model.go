package game

import (
	"fmt"
	"github.com/jesstracy/gapi/db"
)

// note id is provided by default, its the primary key
// for all of our records
type Game struct {
	Id   int    `gorm:"primary_key;"`
	Name string `json:"name"`
}

type GameDLInterface interface {
	RetrieveAllGames() ([]Game, error)
	RetrieveSingleGame(int) (Game, error)
	CreateGame(string) (Game, error)
	DeleteGame(int) error
}

type GameDLGorm struct {
}

func (g *GameDLGorm) RetrieveSingleGame(gameId int) (Game, error) {
	var game Game
	data, err := db.NewDB()
	if err != nil {
		return game, err
	}
	data.Find(&game, gameId)

	return game, nil
}

func (g *GameDLGorm) RetrieveAllGames() ([]Game, error) {
	fmt.Println("geting here")
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	games := []Game{}
	// https://github.com/jinzhu/gorm/blob/master/main.go#L309
	retDB := data.Find(&games) //finding zero.
	fmt.Printf("%+v\n", retDB)
	fmt.Println("Num games:", len(games))
	fmt.Println("Value:", retDB.Value)
	// IDK ¯\_(ツ)_/¯ start with create next time??
	return games, nil
}

func (g *GameDLGorm) CreateGame(gameName string) (Game, error) {
	var game Game
	data, err := db.NewDB()
	if err != nil {
		return game, err
	}
	game = Game{Name: gameName}
	data.Create(&game)
	return game, err
}

func (g *GameDLGorm) DeleteGame(gameId int) error {
	data, err := db.NewDB()
	game := Game{Id: gameId}
	if err != nil {
		return err
	}
	data.Delete(&game)
	return err
}
