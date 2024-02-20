package v2_5

type CurrentWeatherRequest struct {
	//https://api.openweathermap.org/data/3.0/onecall?lat=33.44&lon=-94.04&appid={API key}
	Lat    float64 `json:"lat"`
	Lon    float64 `json:"lon"`
	APIKey string
	Units  string
	Lang   string
}

func NewCurrentWeatherRequest(
	lat float64,
	lon float64,
	APIKey string,
	Lang string,
	Units string,
) *CurrentWeatherRequest {
	return &CurrentWeatherRequest{Lat: lat, Lon: lon, APIKey: APIKey, Lang: Lang, Units: Units}
}

func NewCurrentWeatherResponse() *CurrentWeatherResponse {
	return &CurrentWeatherResponse{}
}

type CurrentWeather struct {
	CurrentWeatherRequest
	CurrentWeatherResponse
}

func NewCurrentWeather(currentWeatherRequest CurrentWeatherRequest) *CurrentWeather {
	return &CurrentWeather{CurrentWeatherRequest: currentWeatherRequest}
}

type CurrentWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		H float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone              int    `json:"timezone"`
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Cod                   int    `json:"cod"`
	TemperatureAssessment string
}
