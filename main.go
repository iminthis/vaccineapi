package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Stats struct {
	DosesGiven             string `json:"dosesGiven"`
	NewDosesGiven          string `json:"newDosesGiven"`
	FullyVaccinated        string `json:"fullyVaccinated"`
	PercentFullyVaccinated string `json:"percentFullyVaccinated"`
}

type Articles []Stats

type Data struct {
	Country string `json:"country"`
	IsoCode string `json:"iso_code"`
	Data    []struct {
		Date                        string
		TotalVaccinations           int     `json:"total_vaccinations"`
		PeopleVaccinated            int     `json:"people_vaccinated"`
		TotalVaccinationsPerHundred float32 `json:"total_vaccinations_per_hundred"`
		PeopleVaccinatedPerHundred  float32 `json:"people_vaccinated_per_hundred"`
		DailyVaccinationsperMillion int     `json:"daily_vaccinations_per_million"`
	}
}

func reqData() {

	resp, err := http.Get("https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/vaccinations.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	/*var datas []Data
	json.Unmarshal([]byte(body), &datas)
	fmt.Printf("Datas : %+v", datas)*/

}

func returnStats(w http.ResponseWriter, r *http.Request) {
	stats := Articles{
		Stats{DosesGiven: "1024", NewDosesGiven: "1024", FullyVaccinated: "1024", PercentFullyVaccinated: "1024"},
	}

	fmt.Println("Endpoint Hit: All Stats Endpoint")
	json.NewEncoder(w).Encode(stats)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", returnStats)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	//handleRequests()
	reqData()
}