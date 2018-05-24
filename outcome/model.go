package outcome

import (
	"github.com/jesstracy/gapi/db"
)

// todo: foreign keys
// todo: include a int score
// todo: include a status win/lose/tie

type Outcome struct {
	Id       int    `gorm:"primary_key;"`
	GameId   int    `json:"gameid"`
	PlayerId int    `json:"playerid"`
	Result   string `json:"result"` // win, loss, tie
	Date     string `json:"date"`   // 2017-08-24 ?
	Score    int    `json:"score"`
}

type OutcomeDLInterface interface {
	RetrieveAllOutcomes() ([]Outcome, error)
	RetrieveSingleOutcome(int) (Outcome, error)
	CreateOutcome(int, int, string, string, int) (*Outcome, error) // TODO just pass in object?
	DeleteOutcome(int) error
}

type OutcomeDLGorm struct {
}

func (o *OutcomeDLGorm) RetrieveSingleOutcome(outcomeId int) (Outcome, error) {
	var outcome Outcome
	data, err := db.NewDB()
	if err != nil {
		return outcome, err
	}
	data.Find(&outcome, outcomeId)

	return outcome, nil
}

func (o *OutcomeDLGorm) RetrieveAllOutcomes() ([]Outcome, error) {
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	outcomes := []Outcome{}
	data.Find(&outcomes)
	return outcomes, nil
}

func (o *OutcomeDLGorm) CreateOutcome(gameId int, playerId int, result string, date string, score int) (*Outcome, error) {
	var outcome Outcome
	data, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	outcome = Outcome{
		GameId:   gameId,
		PlayerId: playerId,
		Result:   result,
		Date:     date,
		Score:    score,
	}
	data.Create(&outcome)
	return &outcome, nil
}

func (o *OutcomeDLGorm) DeleteOutcome(outcomeId int) error {
	data, err := db.NewDB()
	if err != nil {
		return err
	}
	outcome := Outcome{Id: outcomeId}
	data.Delete(&outcome)
	return nil
}
