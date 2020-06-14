package types

// US represents us data structure for covicapi
type US struct {
	Date                     int     `json:"date"`
	DateChecked              string  `json:"dateChecked"`
	Death                    *int    `json:"death"`
	DeathIncrease            *int    `json:"deathIncrease"` // deprecated
	Hash                     string  `json:"hash"`
	Hospitalized             *int    `json:"hospitalized"` // deprecated
	HospitalizedCumulative   *int    `json:"hospitalizedCumulative"`
	HospitalizedCurrently    *int    `json:"hospitalizedCurrently"`
	HospitalizedIncrease     *int    `json:"hospitalizedIncrease"` // deprecated
	InIcuCurrently           *int    `json:"inIcuCurrently"`
	InIcuCumulative          *int    `json:"inIcuCumulative"`
	LastModified             string  `json:"lastModified"` //deprecated
	Negative                 *int    `json:"negative"`
	NegativeIncrease         *int    `json:"negativeIncrease"`
	OnVentilatorCurrently    *int    `json:"onVentilatorCurrently"`
	OnVentilatorCumulative   *int    `json:"onVentilatorCumulative"`
	Pending                  *int    `json:"pending"`
	PosNeg                   *int    `json:"posNeg"`
	Positive                 *int    `json:"positive"`
	PositiveIncrease         *int    `json:"positiveIncrease"`
	Recovered                *int    `json:"recovered"`
	States                   int     `json:"states"`
	Total                    *int    `json:"total"`
	TotalTestResults         *int    `json:"totalTestResults"`
	TotalTestResultsIncrease *int    `json:"totalTestResultsIncrease"`
	PositiveAvg              float64 `json:"positiveAvg"`
	TestsAvg                 float64 `json:"testsAvg"`
	DeathsAvg                float64 `json:"deathsAvg"`
}

type State struct {
	CheckTimeEt              string  `json:"checkTimeEt"`
	CommercialScore          *int    `json:"commercialScore"`
	DataQualityGrade         string  `json:"dataQualityGrade"`
	Date                     int     `json:"date"`
	DateChecked              string  `json:"dateChecked"`
	DateModified             string  `json:"dateModified"`
	Death                    *int    `json:"death"`
	DeathIncrease            *int    `json:"deathIncrease"`
	Fips                     string  `json:"fips"`
	Grade                    string  `json:"grade"`
	Hash                     string  `json:"hash"`
	Hospitalized             *int    `json:"hospitalized"`
	HospitalizedCumulative   *int    `json:"hospitalizedCumulative"`
	HospitalizedCurrently    *int    `json:"hospitalizedCurrently"`
	HospitalizedIncrease     *int    `json:"hospitalizedIncrease"`
	InIcuCurrently           *int    `json:"inIcuCurrently"`
	InIcuCumulative          *int    `json:"inIcuCumulative"`
	LastUpdateEt             string  `json:"lastUpdateEt"`
	Negative                 *int    `json:"negative"`
	NegativeIncrease         *int    `json:"negativeIncrease"`
	NegativeRegularScore     *int    `json:"negativeRegularScore"`
	NegativeScore            *int    `json:"negativeScore"`
	NegativeTestsViral       *int    `json:"negativeTestsViral"`
	OnVentilatorCumulative   *int    `json:"onVentilatorCumulative"`
	OnVentilatorCurrently    *int    `json:"onVentilatorCurrently"`
	Pending                  *int    `json:"pending"`
	PosNeg                   *int    `json:"posNeg"`
	Positive                 *int    `json:"positive"`
	PositiveCasesViral       *int    `json:"positiveCasesViral"`
	PositiveIncrease         *int    `json:"positiveIncrease"`
	PositiveScore            *int    `json:"positiveScore"`
	PositiveTestsViral       *int    `json:"positiveTestsViral"`
	Recovered                *int    `json:"recovered"`
	Score                    *int    `json:"score"`
	State                    string  `json:"state"`
	Total                    *int    `json:"total"`
	TotalTestResults         *int    `json:"totalTestResults"`
	TotalTestResultsIncrease *int    `json:"totalTestResultsIncrease"`
	TotalTestsViral          *int    `json:"totalTestsViral"`
	PositiveAvg              float64 `json:"positiveAvg"`
	TestsAvg                 float64 `json:"testsAvg"`
	DeathsAvg                float64 `json:"deathsAvg"`
}
