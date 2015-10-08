package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Configution type to unmarshal to
type Configuration struct {
	Datasources []string     `json:"data sources"`
	Airports    []Airport    `json:"metros"`
	FlightPaths []FlightPath `json:"routes"`
}

// Save to Disk
func save(args []string) {
	var config Configuration
	config.Datasources = network.Datasources
	config.Airports = make([]Airport, 0, len(network.Airports))
	for _, airport := range network.Airports {
		config.Airports = append(config.Airports, airport)
	}
	config.FlightPaths = network.FlightPaths
	j, _ := json.Marshal(config)
	ioutil.WriteFile("res/save.json", j, 0644)
}

// Load from Disk
func load(args []string) {
	path := fmt.Sprintf("res/%s.json", args[0])
	network = NewNetwork(path)
}

// Help
func help(args []string) {
	commands := make([]string, 0, len(commandMap))
	for command, _ := range commandMap {
		commands = append(commands, command)
	}
	sort.Strings(commands)

	fmt.Println(strings.Join(commands, "\n"))
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
		"save":              save,
		"load":              load,
		"stats":             routeStatistics,
		"exit":              exit,
		"help":              help,
	}

	network = NewNetwork("res/map_data.json")
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
