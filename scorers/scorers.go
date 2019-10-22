package Scorers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Player struct {
	Count   int `json:"count"`
	Filters struct {
		Limit int `json:"limit"`
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	Scorers []struct {
		Player struct {
			ID             int         `json:"id"`
			Name           string      `json:"name"`
			FirstName      string      `json:"firstName"`
			LastName       interface{} `json:"lastName"`
			DateOfBirth    string      `json:"dateOfBirth"`
			CountryOfBirth string      `json:"countryOfBirth"`
			Nationality    string      `json:"nationality"`
			Position       string      `json:"position"`
			ShirtNumber    int         `json:"shirtNumber"`
			LastUpdated    time.Time   `json:"lastUpdated"`
		} `json:"player"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		NumberOfGoals int `json:"numberOfGoals"`
	} `json:"scorers"`
}

func Scorers(league, token string) {

	dataScorer := Player{}
	_ = json.Unmarshal(getData(league, token), &dataScorer)

	var tableDisplay [][]string

	for i := range dataScorer.Scorers {
		goalNumber := dataScorer.Scorers[i].NumberOfGoals
		name := dataScorer.Scorers[i].Player.Name
		team := dataScorer.Scorers[i].Team.Name
		line := []string{strconv.Itoa(goalNumber), name, team}
		tableDisplay = append(tableDisplay, line)
	}

	printTable(tableDisplay)
}

func getData(league, Token string) []byte {
	url := "https://api.football-data.org/v2/competitions/" + league + "/scorers"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Auth-Token", Token)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	return body
}

func printTable(tableDispay [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"GOAL", "NAME", "TEAM"})
	table.SetAlignment(tablewriter.ALIGN_CENTER) // Set Alignment
	table.SetRowLine(true)
	table.SetBorders(tablewriter.Border{Left: false, Top: true, Right: false, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor, tablewriter.FgWhiteColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.Bold})

	for _, v := range tableDispay {
		table.Append(v)
	}
	table.Render()
}
