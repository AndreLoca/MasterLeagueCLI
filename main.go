package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	MatchDay "livesScore/matchDay"
	Scorers "livesScore/scorers"
	"os"
)

// League -> list of supported league
type League struct {
	League []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"league"`
}

// Token attivazione API
const Token string = "YOUR_KEY"

func main() {

	args := os.Args[1:]

	switch args[0] {
	case "-ll":
		printLeague()
	case "-r":
		fmt.Println("not implement yet :)")
	case "-s":
		Scorers.Scorers(args[1], Token)
	case "-md":
		completeURL := "competitions/" + args[1] + "/matches"
		MatchDay.CurrentMatchDay(completeURL, Token)
	case "-m":
		completeURL := "competitions/" + args[1] + "/matches"
		MatchDay.MatchDay(completeURL, args[2], Token)
	case "-la":
		completeURL := "/matches?status=LIVE"
		MatchDay.Live(completeURL, Token)
	case "-l":
		completeURL := "/competitions/" + args[1] + "/matches?status=LIVE"
		MatchDay.Live(completeURL, Token)
	default:
		help()
	}

}

func help() {
	fmt.Println("This is the help page for the MasterLeagueCLI")
	fmt.Println("Use:")
	fmt.Println("\t-ll to see all the aviable leagues")
	fmt.Println("\t-r <league_code> to see the ranking of the league")
	fmt.Println("\t-s <league_code> to see the best scorer of the selected league")
	fmt.Println("\t-md <league_code> to see the result of the current matchday")
	fmt.Println("\t-m <league_code> <number_of_matchday> to see the result of the matchday")
	fmt.Println("\t-la to see all the live match result")
	fmt.Println("\t-l <league_code> to see all the live match of a league")
}

func printLeague() {

	jsonFile, err := os.Open("league.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteLeague, _ := ioutil.ReadAll(jsonFile)
	var listLeague League

	json.Unmarshal(byteLeague, &listLeague)

	fmt.Println()
	for i := 0; i < len(listLeague.League); i++ {
		fmt.Println(listLeague.League[i].Name, " - ", listLeague.League[i].Code)
	}
	fmt.Println()
	fmt.Println("Choose the league that you prefer :)")
}
