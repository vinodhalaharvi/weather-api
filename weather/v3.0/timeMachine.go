package v30

type TimeMachineRequest struct {
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	Dt     int     `json:"dt"`
	APIKey string
	Lang   string
	Units  string
}

func NewTimeMachineRequest(
	lat float64,
	lon float64,
	APIKey string,
) *TimeMachineRequest {
	return &TimeMachineRequest{Lat: lat, Lon: lon, APIKey: APIKey}
}

func NewTimeMachineResponse() *TimeMachineResponse {
	return &TimeMachineResponse{}
}

type TimeMachine struct {
	TimeMachineRequest
	TimeMachineResponse
}

func NewTimeMachine(currentWeatherRequest TimeMachineRequest) *TimeMachine {
	return &TimeMachine{TimeMachineRequest: currentWeatherRequest}
}

type TimeMachineResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Data           []struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Snow struct {
			H float64 `json:"1h"`
		} `json:"snow"`
	} `json:"data"`
}
