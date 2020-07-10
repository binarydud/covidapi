package types

// US represents us data structure for covicapi
type US struct {
	Date                     int     `json:"date" jsonapi:"attr,date"`
	DateChecked              string  `json:"dateChecked" jsonapi:"attr,dateCheckd"`
	Death                    *int    `json:"death" jsonapi:"attr"`
	DeathIncrease            *int    `json:"deathIncrease" jsonapi:"attr"` // deprecated
	Hash                     string  `json:"hash" jsonapi:"primary,national"`
	Hospitalized             *int    `json:"hospitalized" jsonapi:"attr"` // deprecated
	HospitalizedCumulative   *int    `json:"hospitalizedCumulative" jsonapi:"attr"`
	HospitalizedCurrently    *int    `json:"hospitalizedCurrently" jsonapi:"attr"`
	HospitalizedIncrease     *int    `json:"hospitalizedIncrease" jsonapi:"attr"` // deprecated
	InIcuCurrently           *int    `json:"inIcuCurrently" jsonapi:"attr"`
	InIcuCumulative          *int    `json:"inIcuCumulative" jsonapi:"attr"`
	LastModified             string  `json:"lastModified" jsonapi:"attr"` //deprecated
	Negative                 *int    `json:"negative" jsonapi:"attr"`
	NegativeIncrease         *int    `json:"negativeIncrease" jsonapi:"attr"`
	OnVentilatorCurrently    *int    `json:"onVentilatorCurrently" jsonapi:"attr"`
	OnVentilatorCumulative   *int    `json:"onVentilatorCumulative" jsonapi:"attr"`
	Pending                  *int    `json:"pending" jsonapi:"attr"`
	PosNeg                   *int    `json:"posNeg" jsonapi:"attr"`
	Positive                 *int    `json:"positive" jsonapi:"attr"`
	PositiveIncrease         *int    `json:"positiveIncrease" jsonapi:"attr"`
	Recovered                *int    `json:"recovered" jsonapi:"attr"`
	States                   int     `json:"states" jsonapi:"attr"`
	Total                    *int    `json:"total" jsonapi:"attr"`
	TotalTestResults         *int    `json:"totalTestResults" jsonapi:"attr"`
	TotalTestResultsIncrease *int    `json:"totalTestResultsIncrease" jsonapi:"attr"`
	PositiveAvg              float64 `json:"positiveAvg" jsonapi:"attr"`
	TestsAvg                 float64 `json:"testsAvg" jsonapi:"attr"`
	DeathsAvg                float64 `json:"deathsAvg" jsonapi:"attr"`
	PercentagePositive       float64 `json:"percentagePositive" jsonapi:"attr,percentagePositive"`
}

type State struct {
	CheckTimeEt              string  `json:"checkTimeEt" jsonapi:"attr"`
	CommercialScore          *int    `json:"commercialScore" jsonapi:"attr"`
	DataQualityGrade         string  `json:"dataQualityGrade" jsonapi:"attr"`
	Date                     int     `json:"date" jsonapi:"attr"`
	DateChecked              string  `json:"dateChecked" jsonapi:"attr"`
	DateModified             string  `json:"dateModified" jsonapi:"attr"`
	Death                    *int    `json:"death" jsonapi:"attr"`
	DeathIncrease            *int    `json:"deathIncrease" jsonapi:"attr"`
	Fips                     string  `json:"fips" jsonapi:"attr"`
	Grade                    string  `json:"grade" jsonapi:"attr"`
	Hash                     string  `json:"hash" jsonapi:"primary,states"`
	Hospitalized             *int    `json:"hospitalized" jsonapi:"attr"`
	HospitalizedCumulative   *int    `json:"hospitalizedCumulative" jsonapi:"attr"`
	HospitalizedCurrently    *int    `json:"hospitalizedCurrently" jsonapi:"attr"`
	HospitalizedIncrease     *int    `json:"hospitalizedIncrease" jsonapi:"attr"`
	InIcuCurrently           *int    `json:"inIcuCurrently" jsonapi:"attr"`
	InIcuCumulative          *int    `json:"inIcuCumulative" jsonapi:"attr"`
	LastUpdateEt             string  `json:"lastUpdateEt" jsonapi:"attr"`
	Negative                 *int    `json:"negative" jsonapi:"attr"`
	NegativeIncrease         *int    `json:"negativeIncrease" jsonapi:"attr"`
	NegativeRegularScore     *int    `json:"negativeRegularScore" jsonapi:"attr"`
	NegativeScore            *int    `json:"negativeScore" jsonapi:"attr"`
	NegativeTestsViral       *int    `json:"negativeTestsViral" jsonapi:"attr"`
	OnVentilatorCumulative   *int    `json:"onVentilatorCumulative" jsonapi:"attr"`
	OnVentilatorCurrently    *int    `json:"onVentilatorCurrently" jsonapi:"attr"`
	Pending                  *int    `json:"pending" jsonapi:"attr"`
	PosNeg                   *int    `json:"posNeg" jsonapi:"attr"`
	Positive                 *int    `json:"positive" jsonapi:"attr"`
	PositiveCasesViral       *int    `json:"positiveCasesViral" jsonapi:"attr"`
	PositiveIncrease         *int    `json:"positiveIncrease" jsonapi:"attr"`
	PositiveScore            *int    `json:"positiveScore" jsonapi:"attr"`
	PositiveTestsViral       *int    `json:"positiveTestsViral" jsonapi:"attr"`
	Recovered                *int    `json:"recovered" jsonapi:"attr"`
	Score                    *int    `json:"score" jsonapi:"attr"`
	State                    string  `json:"state" jsonapi:"attr"`
	Total                    *int    `json:"total" jsonapi:"attr"`
	TotalTestResults         *int    `json:"totalTestResults" jsonapi:"attr"`
	TotalTestResultsIncrease *int    `json:"totalTestResultsIncrease" jsonapi:"attr"`
	TotalTestsViral          *int    `json:"totalTestsViral" jsonapi:"attr"`
	PositiveAvg              float64 `json:"positiveAvg" jsonapi:"attr"`
	TestsAvg                 float64 `json:"testsAvg" jsonapi:"attr"`
	DeathsAvg                float64 `json:"deathsAvg" jsonapi:"attr"`
	PercentagePositive       float64 `json:"percentagePositive" jsonapi:"attr,percentagePositive"`
}
