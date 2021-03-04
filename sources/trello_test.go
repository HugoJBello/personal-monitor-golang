package sources

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func init(){
	godotenv.Load("../.env")

}

func TestTrelloSource_Extract(t *testing.T) {
	key := os.Getenv("TRELLOKEY")
	token := os.Getenv("TRELLOTOKEN")
	list := os.Getenv("TRELLODEFAULTLIST")

	trello := TrelloSource{
		TrelloKey:   key,
		TrelloToken: token,
		DefaultList: list,
	}
	fmt.Println(key, token,list)

	result,err := trello.Extract()
	assert.NotNil(t, result)
	assert.NotNil(t, result.Cards[0])
	fmt.Println(result)
	assert.Nil(t, err)

}