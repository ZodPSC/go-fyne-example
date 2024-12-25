package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := getWeather()
	if err != nil {
		return
	}
	fmt.Printf("temperature in Moscow=%.2f", res.Main.Temp)
}

type WeatherData struct {
	//Weather []struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	//} `json:"weather"`
}

// func getWeather(result interface{}) error {
func getWeather() (WeatherData, error) {
	const appid = "<your appid>"
	const city = "moscow"
	var url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather"+
		"?q=%s"+
		"&appid=%s", city, appid)
	println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	//return json.NewDecoder(response.Body).Decode(result)

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Unmarshal the JSON data
	weatherData := WeatherData{}

	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return weatherData, nil
}
