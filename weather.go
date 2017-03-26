package main

import (
        "encoding/json"
	"fmt"
	"log"
        "os"
	"net/http"
	"net/url"
)

// https://mholt.github.io/json-to-go/  converts JSON to struct

type CurrentWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID int `json:"id"`
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp float64 `json:"temp"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg int `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt int `json:"dt"`
	Sys struct {
		Type int `json:"type"`
		ID int `json:"id"`
		Message float64 `json:"message"`
		Country string `json:"country"`
		Sunrise int `json:"sunrise"`
		Sunset int `json:"sunset"`
	} `json:"sys"`
	ID int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
}

func ConverToCel(t float64) float64 {
	return t - 273.15 
}

func Temp(UrlBase string,location string, token string) error{
  safeLocation := url.QueryEscape(location)
  safeToken := url.QueryEscape(token)
 
  url := fmt.Sprintf("%s/weather?q=%s&appid=%s",UrlBase,safeLocation, safeToken)
  
  req, err := http.NewRequest("GET", url, nil)

  client := &http.Client{}
 
  resp, err := client.Do(req)

  defer resp.Body.Close()
  
  var cw CurrentWeather

  if err := json.NewDecoder(resp.Body).Decode(&cw); err != nil {
		log.Println(err)
	} 
  fmt.Println("Temp in ",location," ==  ", ConverToCel(cw.Main.Temp))
  return err
}

func main() {
  location := os.Getenv("W_LOCATION")

  if location == "" {
    location = "Valencia,es"
  }
  token := os.Getenv("W_TOKEN")

  Temp("http://api.openweathermap.org/data/2.5/", location, token)
}

