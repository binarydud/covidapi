package types

// US represents us data structure for covicapi
type US struct {
	Date                     int     `json:"date"`
	States                   int     `json:"states"`
	Positive                 int     `json:"positive"`
	Negative                 *int    `json:"negative"`
	Pending                  int     `json:"pending"`
	HospitalizedCurrently    *int    `json:"hospitalizedCurrently"`
	HospitalizedCumulative   *int    `json:"hospitalizedCumulative"`
	InIcuCurrently           int     `json:"inIcuCurrently"`
	InIcuCumulative          int     `json:"inIcuCumulative"`
	OnVentilatorCurrently    int     `json:"onVentilatorCurrently"`
	OnVentilatorCumulative   int     `json:"onVentilatorCumulative"`
	Recovered                int     `json:"recovered"`
	DateChecked              string  `json:"dateChecked"`
	Death                    *int    `json:"death"`
	Hospitalized             *int    `json:"hospitalized"` // deprecated
	LastModified             string  `json:"lastModified"` //deprecated
	Total                    int     `json:"total"`
	TotalTestResults         int     `json:"totalTestResults"`
	PosNeg                   int     `json:"posNeg"`
	DeathIncrease            *int    `json:"deathIncrease"`        // deprecated
	HospitalizedIncrease     *int    `json:"hospitalizedIncrease"` // deprecated
	NegativeIncrease         int     `json:"negativeIncrease"`
	PositiveIncrease         int     `json:"positiveIncrease"`
	TotalTestResultsIncrease int     `json:"totalTestResultsIncrease"`
	Hash                     string  `json:"hash"`
	PositiveAvg              float64 `json:"positiveAvg"`
	TestsAvg                 float64 `json:"positiveAvg"`
	DeathsAvg                float64 `json:"positiveAvg"`
}

type State struct {
	Date                     int     `json:"date"`
	State                    string  `json:"state"`
	Positive                 int     `json:"positive"`
	Negative                 int     `json:"negative"`
	Pending                  *int    `json:"pending"`
	HospitalizedCurrently    int     `json:"hospitalizedCurrently"`
	HospitalizedCumulative   *int    `json:"hospitalizedCumulative"`
	InIcuCurrently           *int    `json:"inIcuCurrently"`
	InIcuCumulative          *int    `json:"inIcuCumulative"`
	OnVentilatorCurrently    *int    `json:"onVentilatorCurrently"`
	OnVentilatorCumulative   *int    `json:"onVentilatorCumulative"`
	Recovered                *int    `json:"recovered"`
	DataQualityGrade         string  `json:"dataQualityGrade"`
	LastUpdateEt             string  `json:"lastUpdateEt"`
	DateModified             string  `json:"dateModified"`
	CheckTimeEt              string  `json:"checkTimeEt"`
	Death                    int     `json:"death"`
	Hospitalized             *int    `json:"hospitalized"`
	DateChecked              string  `json:"dateChecked"`
	TotalTestsViral          int     `json:"totalTestsViral"`
	PositiveTestsViral       *int    `json:"positiveTestsViral"`
	NegativeTestsViral       *int    `json:"negativeTestsViral"`
	PositiveCasesViral       *int    `json:"positiveCasesViral"`
	Fips                     string  `json:"fips"`
	PositiveIncrease         int     `json:"positiveIncrease"`
	NegativeIncrease         int     `json:"negativeIncrease"`
	Total                    int     `json:"total"`
	TotalTestResults         int     `json:"totalTestResults"`
	TotalTestResultsIncrease int     `json:"totalTestResultsIncrease"`
	PosNeg                   int     `json:"posNeg"`
	DeathIncrease            int     `json:"deathIncrease"`
	HospitalizedIncrease     int     `json:"hospitalizedIncrease"`
	Hash                     string  `json:"hash"`
	CommercialScore          int     `json:"commercialScore"`
	NegativeRegularScore     int     `json:"negativeRegularScore"`
	NegativeScore            int     `json:"negativeScore"`
	PositiveScore            int     `json:"positiveScore"`
	Score                    int     `json:"score"`
	Grade                    string  `json:"grade"`
	PositiveAvg              float64 `json:"positiveAvg"`
	TestsAvg                 float64 `json:"positiveAvg"`
	DeathsAvg                float64 `json:"positiveAvg"`
}
