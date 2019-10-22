package MatchDay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Matches struct {
	Competition struct {
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Code        string `json:"code"`
		ID          int    `json:"id"`
		LastUpdated string `json:"lastUpdated"`
		Name        string `json:"name"`
		Plan        string `json:"plan"`
	} `json:"competition"`
	Count   int      `json:"count"`
	Filters struct{} `json:"filters"`
	Matches []struct {
		AwayTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"awayTeam"`
		Group    string `json:"group"`
		HomeTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"homeTeam"`
		ID          int    `json:"id"`
		LastUpdated string `json:"lastUpdated"`
		Matchday    int    `json:"matchday"`
		Referees    []struct {
			ID          int         `json:"id"`
			Name        string      `json:"name"`
			Nationality interface{} `json:"nationality"`
		} `json:"referees"`
		Score struct {
			Duration  string `json:"duration"`
			ExtraTime struct {
				AwayTeam interface{} `json:"awayTeam"`
				HomeTeam interface{} `json:"homeTeam"`
			} `json:"extraTime"`
			FullTime struct {
				AwayTeam int `json:"awayTeam"`
				HomeTeam int `json:"homeTeam"`
			} `json:"fullTime"`
			HalfTime struct {
				AwayTeam int `json:"awayTeam"`
				HomeTeam int `json:"homeTeam"`
			} `json:"halfTime"`
			Penalties struct {
				AwayTeam interface{} `json:"awayTeam"`
				HomeTeam interface{} `json:"homeTeam"`
			} `json:"penalties"`
			Winner string `json:"winner"`
		} `json:"score"`
		Season struct {
			CurrentMatchday int    `json:"currentMatchday"`
			EndDate         string `json:"endDate"`
			ID              int    `json:"id"`
			StartDate       string `json:"startDate"`
		} `json:"season"`
		Stage   string `json:"stage"`
		Status  string `json:"status"`
		UtcDate string `json:"utcDate"`
	} `json:"matches"`
}

func CurrentMatchDay(league, token string) {

	dataMatch := Matches{}
	_ = json.Unmarshal(getData(league, token), &dataMatch)

	var tableDisplay [][]string
	for i := range dataMatch.Matches {
		if dataMatch.Matches[i].Matchday == dataMatch.Matches[i].Season.CurrentMatchday {
			homeTeam := dataMatch.Matches[i].HomeTeam.Name
			awayTeam := dataMatch.Matches[i].AwayTeam.Name
			score := fmt.Sprintf("%d - %d", dataMatch.Matches[i].Score.FullTime.HomeTeam, dataMatch.Matches[i].Score.FullTime.AwayTeam)
			finished := dataMatch.Matches[i].Status
			line := []string{homeTeam, score, awayTeam, finished}
			tableDisplay = append(tableDisplay, line)
		}

	}

	printTable(tableDisplay)

}

func MatchDay(league, matchday, token string) {

	dataMatch := Matches{}
	_ = json.Unmarshal(getData(league, token), &dataMatch)
	matchRequest, _ := strconv.Atoi(matchday)

	var tableDisplay [][]string
	for i := range dataMatch.Matches {
		if dataMatch.Matches[i].Matchday == matchRequest {
			homeTeam := dataMatch.Matches[i].HomeTeam.Name
			awayTeam := dataMatch.Matches[i].AwayTeam.Name
			score := fmt.Sprintf("%d - %d", dataMatch.Matches[i].Score.FullTime.HomeTeam, dataMatch.Matches[i].Score.FullTime.AwayTeam)
			finished := dataMatch.Matches[i].Status
			line := []string{homeTeam, score, awayTeam, finished}
			tableDisplay = append(tableDisplay, line)
		}

	}

	printTable(tableDisplay)

}

func Live(url, token string) {

	dataMatch := Matches{}
	_ = json.Unmarshal(getData(url, token), &dataMatch)

	var tableDisplay [][]string
	for i := range dataMatch.Matches {
		homeTeam := dataMatch.Matches[i].HomeTeam.Name
		awayTeam := dataMatch.Matches[i].AwayTeam.Name
		score := fmt.Sprintf("%d - %d", dataMatch.Matches[i].Score.FullTime.HomeTeam, dataMatch.Matches[i].Score.FullTime.AwayTeam)
		finished := dataMatch.Matches[i].Status
		line := []string{homeTeam, score, awayTeam, finished}
		tableDisplay = append(tableDisplay, line)

	}

	printTable(tableDisplay)
}

func printTable(tableDispay [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"HOME", "SCORE", "AWAY", "STATUS"})
	table.SetAlignment(tablewriter.ALIGN_CENTER) // Set Alignment
	table.SetRowLine(true)
	table.SetBorders(tablewriter.Border{Left: false, Top: true, Right: false, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold})

	for _, v := range tableDispay {
		table.Append(v)
	}
	table.Render()
}

func getData(competeURL, Token string) []byte {
	url := "https://api.football-data.org/v2/" + competeURL
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", Token)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	return body
}
