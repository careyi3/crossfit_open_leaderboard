package http_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/careyi3/crossfit_open_leaderboard/models"
)

func FetchPage(year int, division int, pageNum int) *models.Response {
	path := fmt.Sprintf("https://c3po.crossfit.com/api/leaderboards/v2/competitions/open/%d/leaderboards?&division=%d&page=%d", year, division, pageNum)
	response, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
	}
	return parse(response)
}

func parse(response *http.Response) *models.Response {
	out := &models.Response{}
	if response.Body != nil {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if body != nil {
			err = json.Unmarshal(body, &out)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	return out
}
