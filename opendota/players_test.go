package opendota

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayersService_Player(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/players/34505203", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"tracked_until":null,"competitive_rank":null,"solo_competitive_rank":null,"mmr_estimate":{"estimate":7932},"profile":{"account_id":34505203,"personaname":"$a$uk3","name":"MinD_ContRoL","cheese":0,"steamid":"76561197994770931","avatar":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg","avatarmedium":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg","avatarfull":"https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg","profileurl":"http://steamcommunity.com/id/MinD_ContRoL/","last_login":null,"loccountrycode":"BG"}}`)
	})

	profile := Profile{
		AccountID:      34505203,
		Personaname:    "$a$uk3",
		Name:           "MinD_ContRoL",
		Cheese:         0,
		Steamid:        "76561197994770931",
		Avatar:         "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221.jpg",
		Avatarmedium:   "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_medium.jpg",
		Avatarfull:     "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/9b/9bb81de39d4a3104fa7e85ce4ed8564082201221_full.jpg",
		Profileurl:     "http://steamcommunity.com/id/MinD_ContRoL/",
		LastLogin:      interface{}(nil),
		Loccountrycode: "BG",
	}

	mmrestimate := MmrEstimate{
		Estimate: 7932,
	}

	expected := &Player{
		TrackedUntil:        interface{}(nil),
		CompetitiveRank:     interface{}(nil),
		SoloCompetitiveRank: interface{}(nil),
		MmrEstimate:         mmrestimate,
		Profile:             profile,
	}

	accountID := &PlayersParam{
		AccountID: 34505203,
	}

	client := NewClient(httpClient)
	player, _, err := client.PlayersService.Player(accountID)
	assert.Nil(t, err)
	assert.Equal(t, expected, player)
}