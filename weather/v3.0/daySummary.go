package v30

type DaySummaryRequest struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Date    string  `json:"date"`
	APIKey  string
	Units   string
	Lang    string
	Exclude string
}

func NewDaySummaryRequest(
	lat float64,
	lon float64,
	APIKey string,
) *DaySummaryRequest {
	return &DaySummaryRequest{Lat: lat, Lon: lon, APIKey: APIKey}
}

func NewDaySummaryResponse() *DaySummaryResponse {
	return &DaySummaryResponse{}
}

type DaySummary struct {
	DaySummaryRequest
	DaySummaryResponse
}

func NewDaySummary(currentWeatherRequest DaySummaryRequest) *DaySummary {
	return &DaySummary{DaySummaryRequest: currentWeatherRequest}
}

type DaySummaryResponse struct {
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Tz         string  `json:"tz"`
	Date       string  `json:"date"`
	Units      string  `json:"units"`
	CloudCover struct {
		Afternoon float64 `json:"afternoon"`
	} `json:"cloud_cover"`
	Humidity struct {
		Afternoon float64 `json:"afternoon"`
	} `json:"humidity"`
	Precipitation struct {
		Total float64 `json:"total"`
	} `json:"precipitation"`
	Temperature struct {
		Min       float64 `json:"min"`
		Max       float64 `json:"max"`
		Afternoon float64 `json:"afternoon"`
		Night     float64 `json:"night"`
		Evening   float64 `json:"evening"`
		Morning   float64 `json:"morning"`
	} `json:"temperature"`
	Pressure struct {
		Afternoon float64 `json:"afternoon"`
	} `json:"pressure"`
	Wind struct {
		Max struct {
			Speed     float64 `json:"speed"`
			Direction float64 `json:"direction"`
		} `json:"max"`
	} `json:"wind"`
}
