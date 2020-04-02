package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tb "github.com/yanzay/tbot/v2"
)

type Corona struct {
	Country             string `json:"country"`
	Cases               int   `json:"cases"`
	TodayCases          int   `json:"todayCases"`
	Deaths              int   `json:"deaths"`
	TodayDeaths         int   `json:"todayDeaths"`
	Recovered           int   `json:"recovered"`
	Active              int   `json:"active"`
	Critical            int   `json:"critical"`
	CasesPerOneMillion  int   `json:"casesPerOneMillion"`
	DeathsPerOneMillion float64   `json:"deathsPerOneMillion"`
	Updated             int64 `json:"updated"`
}



var Token = "1111391792:AAFARj7gRp77r5fkiXm4dAeZyw02PFW1E5I"

type Details struct {
	TotalDeaths         int
	WorldTotalDeath     int
	TotalConfirmed      int
	WorldTotalConfirmed int
	NewDeaths           int
	WorldNewDeaths      int
	NewConfirmed        int
	WorldNewConfirmed   int
	NewRecovered        int
	WorldNewRecovered   int
	TotalRecovered      int
	WorldTotalRecovered int
}

func main() {
	bot := tb.New(Token)
	c := bot.Client()

	bot.HandleMessage("/start", func(m *tb.Message) {
		response := fmt.Sprintf("Hi %s, my name is %s and I will give information about the corona in the country you want!", m.Chat.Username, "COVIDGRAM")
		c.SendMessage(m.Chat.ID, response)
	})

	bot.HandleMessage("/help", func(m *tb.Message) {
		response := fmt.Sprintf("Example -> Turkey, tr, Tr, turkey. World -> /world")
		c.SendMessage(m.Chat.ID, response)
	})

	bot.HandleMessage("/world", func(m *tb.Message) {
		client := &http.Client{}
		url := fmt.Sprintf("https://corona.lmao.ninja/all")
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()
		var corona Corona
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &corona)
		if err != nil {
			log.Fatal(err.Error())
		}

		text1 := fmt.Sprintf("Total confirmed -> %d", corona.Cases)
		text2 := fmt.Sprintf("Total deaths -> %d", corona.Deaths)
		text3 := fmt.Sprintf("Total recovered -> %d", corona.Recovered)		

		response := fmt.Sprintf("%s #StayHome\n%s\n%s\n%s\n", "World", text1, text2, text3)
		c.SendMessage(m.Chat.ID, response)
	})

	bot.HandleMessage("", func(m *tb.Message) {
		client := &http.Client{}
		url := fmt.Sprintf("https://corona.lmao.ninja/countries/%s", m.Text)
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()
		var corona Corona
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &corona)
		if err != nil {
			log.Fatal(err.Error())
		}

		text1 := fmt.Sprintf("New confirmed -> %d", corona.TodayCases)
		text2 := fmt.Sprintf("New deaths -> %d", corona.TodayDeaths)

		text3 := fmt.Sprintf("Total confirmed -> %d", corona.Cases)
		text4 := fmt.Sprintf("Total deaths -> %d", corona.Deaths)
		text5 := fmt.Sprintf("Total recovered -> %d", corona.Recovered)
		

		response := fmt.Sprintf("%s #StayHome\n%s\n%s\n%s\n%s\n%s\n", corona.Country, text1, text2, text3, text4, text5)
		c.SendMessage(m.Chat.ID, response)
	})

	log.Fatal(bot.Start())
}
