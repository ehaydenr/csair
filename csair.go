package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Configution type to unmarshal to
type Configuration struct {
	Datasources []string     `json:"data sources"`
	Airports    []Airport    `json:"metros"`
	FlightPaths []FlightPath `json:"routes"`
}

// Help
func help(args []string) {
	for command, _ := range commandMap {
		fmt.Println(command)
	}
}

// Exit
func exit(args []string) {
	os.Exit(0)
}

// Holds map of commands that CLI offers
var commandMap map[string]func([]string)

// gobal network after it's built
var network *Network

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
		"shortestRoute":     shortestRoute,
		"addCity":           updateCity,
		"removeCity":        removeCity,
		"updateCity":        updateCity,
		"removeRoute":       removeRoute,
		"addRoute":          addRoute,
		"merge":             mergeJSON,
		"exit":              exit,
		"help":              help,
	}

	network = NewNetwork()
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
