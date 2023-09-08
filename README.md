# MWeather
MWeather is a CLI application written in Go (v1.18.1) that prints basic weather information basedon
on location,specified as an argument. It uses Nominatim.OpenStreetMap and Open-Meteo APIs
to get the data.

## Building
```
go build mweather.go
```

## Running
```
mweather <city_name>
```
for example:
```
mweather poznan
```
Mweather takes advanteage of Nominatim's fuzzy search so name of the city does not have to be 
precise.

## Usage
Designed to be run with `.bashrc`, so it prints current info on top of newly created terminals.
