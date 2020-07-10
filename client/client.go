package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"

	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/binarydud/covidapi/types"
)

// HTTPClient for covidtracking.com
type HTTPClient struct {
	URL string
}

// Max returns the larger of x or y
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
func calculateUSPercentagePositive(data []types.US) []types.US {
	newData := make([]types.US, 0)
	for _, i := range data {
		newPositiveCases := *i.PositiveIncrease
		newTests := *i.TotalTestResultsIncrease
		i.PercentagePositive = float64(newPositiveCases / newTests)
		newData = append(newData, i)
	}
	return newData
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
func calculateForStateWindow(data []types.State) types.State {
	window := 7

	positive := movingaverage.New(window)
	tests := movingaverage.New(window)
	deaths := movingaverage.New(window)
	percentPositive := movingaverage.New(window)
	item := data[len(data)-1]
	for _, i := range data {
		newPositiveCases := Max(0, *i.PositiveIncrease)
		newTests := Max(0, *i.TotalTestResultsIncrease)
		positive.Add(float64(newPositiveCases))
		tests.Add(float64(newTests))
		deaths.Add(float64(*i.DeathIncrease))
		percentPositive.Add(float64(newPositiveCases) / float64(newTests) * 100)
	}
	item.PositiveAvg = positive.Avg()
	item.DeathsAvg = deaths.Avg()
	item.TestsAvg = tests.Avg()
	item.PercentagePositive = percentPositive.Avg()
	return item
}

// New creates new http client for covidtracking
func New() *HTTPClient {
	client := &HTTPClient{URL: "https://covidtracking.com"}
	return client
}

// ByStates ...
func (client *HTTPClient) ByStates() ([]types.State, error) {
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

		positive := Max(0, *i.PositiveIncrease)
		deaths := Max(0, *i.DeathIncrease)
		tests := Max(0, *i.TotalTestResultsIncrease)
		i.PositiveIncrease = &positive
		i.DeathIncrease = &deaths
		i.TotalTestResultsIncrease = &tests

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

// ByNational ...
func (client *HTTPClient) ByNational() ([]types.US, error) {
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
	for i, item := range items {
		positive := Max(0, *item.PositiveIncrease)
		deaths := Max(0, *item.DeathIncrease)
		tests := Max(0, *item.TotalTestResultsIncrease)
		items[i].PositiveIncrease = &positive
		items[i].DeathIncrease = &deaths
		items[i].TotalTestResultsIncrease = &tests
	}

	if err != nil {
		return nil, err
	}
	items = calculateUSMovingAverage(items)
	return items, nil
}

// ByState ...
func (client *HTTPClient) ByState(state string) (*types.State, error) {
	url := fmt.Sprintf("%s/api/v1/states/%s/daily.json", client.URL, strings.ToLower(state))
	log.Print(url)
	resp, err := http.Get(url)
	var items []types.State

	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Date < items[j].Date
	})
	index := len(items) - 1
	window := 7
	start := Max(0, index-window)
	end := Min(len(items), index+1)
	previous := items[start:end]
	// lastWeek := items[first:last]

	item := calculateForStateWindow(previous)
	positive := Max(0, *item.PositiveIncrease)
	deaths := Max(0, *item.DeathIncrease)
	tests := Max(0, *item.TotalTestResultsIncrease)
	item.PositiveIncrease = &positive
	item.DeathIncrease = &deaths
	item.TotalTestResultsIncrease = &tests
	return &item, nil
}
