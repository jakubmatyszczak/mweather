package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type API int64
const (
    GeoCode API = 0
    OpenWeather API = 1
)

func callAPI(api API, data map[string]string) ([]byte, error){
    url := ""
    if(api == GeoCode){
        url = "https://nominatim.openstreetmap.org/search?addressdetails=1&format=jsonv2&limit=1&q="
        url = url + data["city"]
    }
    if(api == OpenWeather){
        url = "https://api.open-meteo.com/v1/forecast?&current_weather=true&hourly=temperature_2m,relativehumidity_2m,windspeed_10m"
        url = url + "&longitude=" + data["lon"]
        url = url + "&latitude=" + data["lat"]
    }
    req, err := http.NewRequest("GET", url, nil)
    if(err != nil){
        return nil, err
    }
    response, err := http.DefaultClient.Do(req)
    if(err != nil){
        return nil, err
    }
    defer response.Body.Close()
    responseBody, err := io.ReadAll(response.Body)
    if(err!=nil){
        return nil, err
    }
    if(api == GeoCode){
        responseBody = responseBody[1:len(responseBody)-1]
    }
    return responseBody, nil
}

func getSkyCondFromWeatherCode(code int) string{
    if (code == 0){
        return "Clear"
    }
    if(code <= 3){
        return "Partly cloudy"
    }
    if(code <= 48){
        return "Foggy"
    }
    if(code <= 55){
        return "Drizzle"
    }
    if(code <= 57){
        return "Freezeing Drizzle"
    }
    if(code <= 63){
        return "Rain"
    }
    if(code <= 65){
        return "Heavy Rain"
    }
    if(code <= 67){
        return "Freezing Rain"
    }
    if(code <= 77){
        return "Snow"
    }
    if(code <= 81){
        return "Showers"
    }
    if(code <= 82){
        return "Violent Showers"
    }
    if(code <= 86){
        return "Snow Showers"
    }
    if(code <= 95){
        return "Thunderstorm"
    }
    if(code <= 99){
        return "Thunderstorm with hail"
    }
    return "Unknown"
}

func main (){
    jsonBytes, err := callAPI(GeoCode, map[string]string{"city":"poznan"})
    if(err != nil){
        log.Fatal(err)
    }
    var x any
    json.Unmarshal(jsonBytes, &x)
    if(err != nil){
        log.Fatal(err)
    }
    lat :=  x.(map[string]interface{})["lat"].(string)
    lon :=  x.(map[string]interface{})["lon"].(string)
    cityName :=  "Poznań"
    jsonBytes, err = callAPI(OpenWeather, map[string]string{"lat":lat, "lon":lon})
    if(err != nil){
        log.Fatal(nil)
    }
    json.Unmarshal(jsonBytes, &x)
    current := x.(map[string]interface{})["current_weather"].(map[string]interface{})
    log.Println(current)
    temperature := current["temperature"].(float64)
    windspeed := current["windspeed"].(float64)
    windCategoty := "Breeze"
    skyCondition := getSkyCondFromWeatherCode(int(current["weathercode"].(float64)))

    fmt.Println(cityName, ": ", temperature, "°C, Sky:", skyCondition ,", wind:", windspeed, "km/h (",
        windCategoty, "), ")
}

