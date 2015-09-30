package main

import (
	"fmt"
	"strings"
)

type Network struct {
	Cities      []string
	Airports    map[string]Airport
	FlightPaths []FlightPath
}

// Return list of continents
func (network Network) computeListOfContinents() []string {
	continents := make([]string, 0, 7)
	continentMap := network.computeAirportContinentResidency()
	for continent, _ := range continentMap {
		continents = append(continents, continent)
	}
	return continents
}

// Return flight paths from airport
func (network Network) computeNonstopFlights(code string) ([]FlightPath, error) {
	if _, ok := network.Airports[code]; !ok {
		return nil, NoSuchAirportError("Airport doesn't exist.")
	}

	flightPaths := make([]FlightPath, 0)
	for _, path := range network.FlightPaths {
		if path.Ports[0] == code || path.Ports[1] == code {
			flightPaths = append(flightPaths, path)
		}
	}
	return flightPaths, nil
}

// Return the longest flight in the network
// Returns path
func (network Network) computeLongestFlight() FlightPath {
	var longest FlightPath
	for _, path := range network.FlightPaths {
		if path.Distance > longest.Distance {
			longest = path
		}
	}
	return longest
}

// Compute shortest flight in network
// Returns codes of the two airports
func (network Network) computeShortestFlight() FlightPath {
	var shortest = FlightPath{
		nil,
		MaxInt,
	}
	for _, path := range network.FlightPaths {
		if path.Distance < shortest.Distance {
			shortest = path
		}
	}
	return shortest
}

// Compute average distance of all the flights in the network
// Return average
func (network Network) computeAverageDistance() int {
	sum, denom := 0, 0

	for _, path := range network.FlightPaths {
		sum += path.Distance
		denom++
	}

	return int(float32(sum) / float32(denom))
}

// Compute biggest city by population
// Return code for the airport
func (network Network) computeBiggestCity() string {
	var biggest Airport

	for _, airport := range network.Airports {
		if airport.Population > biggest.Population {
			biggest = airport
		}
	}

	return biggest.Code
}

// Compute smallest city by population
// Return code for the airport
func (network Network) computeSmallestCity() string {
	smallest := Airport{
		Population: MaxInt,
	}

	for _, airport := range network.Airports {
		if airport.Population < smallest.Population {
			smallest = airport
		}
	}

	return smallest.Code
}

// Compute average city population
// Return average
func (network Network) computeAveragePopulation() int {
	sum, denom := 0, 0

	for _, airport := range network.Airports {
		sum += airport.Population
		denom++
	}

	return int(float32(sum) / float32(denom))
}

// Compute list of continents and the cities in them
// Return map of continents to airport codes
func (network Network) computeAirportContinentResidency() map[string][]string {
	continentMap := make(map[string][]string)
	for _, airport := range network.Airports {
		continent := airport.Continent

		if list, ok := continentMap[continent]; !ok {
			newList := []string{airport.Code}
			continentMap[continent] = newList
		} else {
			continentMap[continent] = append(list, airport.Code)
		}

	}
	return continentMap
}

// Compute hub cities - cities with most connections
// Return list of airport codes
func (network Network) computeHubCities() []string {
	ocurrenceMap := make(map[string]int)
	for _, path := range network.FlightPaths {
		ocurrenceMap[path.Ports[0]]++
		ocurrenceMap[path.Ports[1]]++
	}

	max := 0
	for _, count := range ocurrenceMap {
		if max < count {
			max = count
		}
	}

	finalList := make([]string, 0, len(network.Airports))
	for code, count := range ocurrenceMap {
		if count == max {
			finalList = append(finalList, code)
		}
	}

	return finalList
}

// Compute map url
func (network Network) computeMapUrl() string {
	locations := make([]string, len(network.FlightPaths))
	for i, path := range network.FlightPaths {
		locations[i] = fmt.Sprintf("%s-%s", path.Ports[0], path.Ports[1])
	}
	return fmt.Sprintf("%s%s", UrlPrefix, strings.Join(locations, ","))
}
