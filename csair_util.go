package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ls
func getCityList(args []string) {
	fmt.Println(network.computeCityList())
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

// Shortest Route
func shortestRoute(args []string) {
	fmt.Println(network.computeShortestRoute(args[0], args[1]))
}

// Update city
func updateCity(args []string) {
	var airport Airport
	path := fmt.Sprintf("res/%s.json", args[0])
	data, _ := ioutil.ReadFile(path)

	if err := json.Unmarshal(data, &airport); err != nil {
		panic(err)
	}

	network.updateCity(airport)
}

// Remove City
func removeCity(args []string) {
	network.removeCity(args[0])
}

// Remove Route
func removeRoute(args []string) {
	network.removeRoute(args[0], args[1])
}

// Add Route
func addRoute(args []string) {
	var fp FlightPath
	path := fmt.Sprintf("res/%s.json", args[0])
	data, _ := ioutil.ReadFile(path)

	if err := json.Unmarshal(data, &fp); err != nil {
		panic(err)
	}
	network.addRoute(fp)
}

// Merge JSON
func mergeJSON(args []string) {
	var config Configuration
	path := fmt.Sprintf("res/%s.json", args[0])
	data, _ := ioutil.ReadFile(path)

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	network.MergeJSON(config)
}

// Get Route statistics
func routeStatistics(args []string) {
	stats, _ := network.computeRouteStatistics(args)
	fmt.Println(stats)
}
