package client

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"

	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/binarydud/covidapi/types"
)

type HttpClient struct {
	URL string
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func calculateUSMovingAverage(data []types.US) []types.US {
	newData := make([]types.US, 0)
	for index, item := range data {
		window := 7
		start := Max(0, index-window)
		end := Min(len(data), index+1)
		previous := data[start:end]
		positive := movingaverage.New(window)
		tests := movingaverage.New(window)
		deaths := movingaverage.New(window)
		for _, i := range previous {
			positive.Add(float64(*i.PositiveIncrease))
			tests.Add(float64(*i.TotalTestResultsIncrease))
			deaths.Add(float64(*i.DeathIncrease))
		}
		item.PositiveAvg = positive.Avg()
		item.DeathsAvg = deaths.Avg()
		item.TestsAvg = tests.Avg()
		newData = append(newData, item)
	}
	return newData
}
func calculateStateMovingAverage(data []types.State) []types.State {
	newData := make([]types.State, 0)

	for index, item := range data {
		window := 7
		start := Max(0, index-window)
		end := Min(len(data), index+1)
		previous := data[start:end]
		positive := movingaverage.New(window)
		tests := movingaverage.New(window)
		deaths := movingaverage.New(window)
		for _, i := range previous {
			positive.Add(float64(*i.PositiveIncrease))
			tests.Add(float64(*i.TotalTestResultsIncrease))
			deaths.Add(float64(*i.DeathIncrease))
		}
		item.PositiveAvg = positive.Avg()
		item.DeathsAvg = deaths.Avg()
		item.TestsAvg = tests.Avg()
		newData = append(newData, item)
	}
	return newData
}

func NewClient() *HttpClient {
	client := &HttpClient{URL: "https://covidtracking.com"}
	return client
}

type State struct {
	DateChecked  time.Time `json:"dateChecked"`
	DateModified string    `json:"dateModified"`
	Death        int       `json:"death"`
	Hospitalized int       `json:"hospitalized"`
	Negative     int       `json:"negative"`
	Positive     int       `json:"positive"`
	Total        int       `json:"total"`
	Province     string    `json:"state"`
}

func (client *HttpClient) ByStates() ([]types.State, error) {
	// /api/v1/states/daily.json
	url := client.URL + "/api/v1/states/daily.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var items []types.State

	err = json.NewDecoder(resp.Body).Decode(&items)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})
	log.Printf("states %v", len(items))

	m := make(map[string][]types.State)
	for _, i := range items {
		m[i.State] = append(m[i.State], i)
	}

	newStates := make([]types.State, 0)

	for _, values := range m {

		avgs := calculateStateMovingAverage(values)

		newStates = append(newStates, avgs...)
	}

	if err != nil {
		return nil, err
	}
	return newStates, nil
}
func (client *HttpClient) ByNational() ([]types.US, error) {
	///api/v1/us/daily.json
	url := client.URL + "/api/v1/us/daily.json"

	resp, err := http.Get(url)
	if err != nil {

		return nil, err
	}
	var items []types.US

	err = json.NewDecoder(resp.Body).Decode(&items)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})

	if err != nil {
		return nil, err
	}
	items = calculateUSMovingAverage(items)
	return items, nil
}
