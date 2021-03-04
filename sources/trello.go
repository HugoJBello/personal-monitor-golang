package sources

import (
	"encoding/json"
	"fmt"
	"github.com/HugoJBello/personal-monitor-golang/models"
	"github.com/HugoJBello/personal-monitor-golang/utils"
	"io/ioutil"
	"net/http"
)

type TrelloSource struct {
	TrelloKey string
	TrelloToken string
	DefaultList string
}

func (trello *TrelloSource) Extract() (trelloResult *models.TrelloResult, err error){
	result := models.TrelloResult{}
	cards := []models.Card{}

	url := "https://api.trello.com/1/lists/"+trello.DefaultList+ "/cards?key="+trello.TrelloKey+ "&token=" + trello.TrelloToken

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &cards)
	result.Cards = cards

	return &result, nil
}

func (trello *TrelloSource) Display(trelloResult *models.TrelloResult) (rr error) {
	utils.ClearConsole()
	for _, card := range trelloResult.Cards {
		fmt.Println("* " + card.Name + " ---" + card.DateLastActivity.Format("2006-01-02 15:04:05"))
	}
	return nil
}
