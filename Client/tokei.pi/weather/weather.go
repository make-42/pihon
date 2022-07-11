package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

var Temperature float64
var Conditions string
var AQI int
var Components APIComponents
var Sunrise int
var Sunset int

var Latitude float32
var Longitude float32
var CityOfObservation string
var OpenWeatherMapAPIKey string

type APICoords struct {
	Lon float64
	Lat float64
}

type APIWeather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type APIWeatherMain struct {
	Temp      float64
	FeelsLike float64
	TempMin   float64
	TempMax   float64
	Pressure  int
	Humidity  int
}

type APIWind struct {
	Speed float64
	Deg   float64
}

type APIClouds struct {
	All int
}

type APISys struct {
	Type    int
	Id      int
	Country string
	Sunrise int
	Sunset  int
}

type APIWeatherData struct {
	Coord      APICoords
	Weather    []APIWeather
	Base       string
	Main       APIWeatherMain
	Visibility int
	Wind       APIWind
	Clouds     APIClouds
	DT         int
	Sys        APISys
	Timezone   int
	Id         int
	Name       string
	Cod        int
}

type APIComponents struct {
	CO    float64
	NO    float64
	NO2   float64
	O3    float64
	SO2   float64
	PM2_5 float64
	PM10  float64
	NH3   float64
}

type APIPollutionMain struct {
	AQI int
}

type APIPollutionDataObject struct {
	Main       APIPollutionMain
	Components APIComponents
	DT         int
}

type APIPollutionData struct {
	Coord APICoords
	List  []APIPollutionDataObject
}

func UpdateCurrentWeather() {
	// Normal weather info
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", Latitude, Longitude, OpenWeatherMapAPIKey))
	checkError(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var APIWeatherDataParsed APIWeatherData
	json.Unmarshal(body, &APIWeatherDataParsed)
	Temperature = APIWeatherDataParsed.Main.Temp - 273.15
	Conditions = APIWeatherDataParsed.Weather[0].Main
	Sunrise = APIWeatherDataParsed.Sys.Sunrise
	Sunset = APIWeatherDataParsed.Sys.Sunset
	// Air pollution info
	respPollution, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/air_pollution?lat=%f&lon=%f&appid=%s", Latitude, Longitude, OpenWeatherMapAPIKey))
	checkError(err)
	defer respPollution.Body.Close()
	bodyPollution, err := io.ReadAll(respPollution.Body)
	var APIPollutionDataParsed APIPollutionData
	json.Unmarshal(bodyPollution, &APIPollutionDataParsed)
	AQI = APIPollutionDataParsed.List[0].Main.AQI
	Components = APIPollutionDataParsed.List[0].Components
}
