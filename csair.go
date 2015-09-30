package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Configution type to unmarshal to
type Configuration struct {
	Datasources []string     `json:"data sources"`
	Airports    []Airport    `json:"metros"`
	FlightPaths []FlightPath `json:"routes"`
}

// ls
func getCityList(args []string) {
	fmt.Println(network.Cities)
}

// lookup <code>
func lookupAirport(args []string) {
	for _, str := range args {
		if airport, ok := network.Airports[str]; !ok {
			fmt.Printf("Failed to find airport: %s\n", str)
		} else {
			fmt.Println(airport)
			nonstop, _ := network.computeNonstopFlights(airport.Code)
			fmt.Println(nonstop)
		}
	}
}

// longest flight
func longestFlight(args []string) {
	fmt.Println(network.computeLongestFlight())
}

// shortest flight
func shortestFlight(args []string) {
	fmt.Println(network.computeShortestFlight())
}

// average flight distance
func averageFlightDistance(args []string) {
	fmt.Println(network.computeAverageDistance())
}

// biggest city
func biggestCity(args []string) {
	fmt.Println(network.computeBiggestCity())
}

// smallest city
func smallestCity(args []string) {
	fmt.Println(network.computeSmallestCity())
}

// average population
func averagePopulation(args []string) {
	fmt.Println(network.computeAveragePopulation())
}

// continent list
func continentList(args []string) {
	fmt.Println(network.computeListOfContinents())
}

// hubs
func hubList(args []string) {
	fmt.Println(network.computeHubCities())
}

// Map
func mapUrl(args []string) {
	fmt.Println(network.computeMapUrl())
}

// Exit
func exit(args []string) {
	os.Exit(0)
}

// Holds map of commands that CLI offers
var commandMap map[string]func([]string)

// gobal network after it's built
var network Network

// Build command map and network
func init() {
	commandMap = map[string]func([]string){
		"ls":                getCityList,
		"lookup":            lookupAirport,
		"longest":           longestFlight,
		"shortest":          shortestFlight,
		"averageLength":     averageFlightDistance,
		"biggest":           biggestCity,
		"smallest":          smallestCity,
		"averagePopulation": averagePopulation,
		"continents":        continentList,
		"hubs":              hubList,
		"map":               mapUrl,
		"exit":              exit,
	}

	network = buildNetwork()

}

// Build a network from map_data.json
func buildNetwork() Network {
	data, err := ioutil.ReadFile("map_data.json")
	if err != nil {
		panic(err)
	}

	config := Configuration{}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	airportMap := make(map[string]Airport)
	cities := make([]string, 0, 50)
	for _, airport := range config.Airports {
		airportMap[airport.Code] = airport
		cities = append(cities, airport.Code)
	}

	// Build network
	network := Network{
		cities,
		airportMap,
		config.FlightPaths,
	}
	return network
}

// Main
// Read commands from stdin and call appropriate functions
func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		if line, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				return
			} else {
				panic(err)
			}
		} else {
			line = strings.TrimSpace(line)
			args := strings.Split(line, " ")
			if fn, ok := commandMap[args[0]]; !ok {
				fmt.Println("Invalid Command.")
			} else {
				fn(args[1:])
			}
		}
	}
}
