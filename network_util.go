package main

import (
	"errors"
	"fmt"
	graphlib "github.com/ehaydenr/algorithms/graph"
	"math"
	"strings"
)

type RouteStats struct {
	Distance int
	Cost     float32
	Time     float32
}

// Compute shortest route between two cities
func (network Network) computeShortestRoute(origin, destination string) []string {
	v1 := network.nodeMap[origin]
	v2 := network.nodeMap[destination]
	network.graph.Dijkstra(v1.Id, v2.Id)

	var prev *graphlib.Vertex = v2.Prev

	result := make([]string, 0)
	result = append(result, v2.Value.(string))

	for prev != nil {
		result = append(result, prev.Value.(string))
		prev = prev.Prev
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// Compute City List
func (network Network) computeCityList() []string {
	cities := make([]string, 0, len(network.Airports))
	for _, city := range network.Airports {
		cities = append(cities, city.Name)
	}
	return cities
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

// Compute Route Statistics
func (network Network) computeRouteStatistics(args []string) (RouteStats, error) {
	dist, cost, time := 0, float32(0), float32(0)
	costMultiplier := float32(0.35)

	for i := 0; i < len(args)-1; i++ {

		// Distance
		dist_t, err := network.computeDistance(args[i], args[i+1])
		if err != nil {
			return RouteStats{}, err
		}
		dist += dist_t

		// Time
		time_t := network.computeDistTime(dist_t)
		if err != nil {
			return RouteStats{}, err
		}
		time += time_t

		if i != 0 {
			outbound, _ := network.computeNonstopFlights(args[i])
			layover := float32(2) - float32(len(outbound))*float32(1.0/6.0)
			time += layover
		}

		// Cost
		cost += costMultiplier * float32(dist_t)
		if costMultiplier > 0 {
			costMultiplier -= 0.05
		}
	}
	return RouteStats{dist, cost, time}, nil
}

// Get distance between cities
func (network Network) computeDistance(c1, c2 string) (int, error) {
	for _, neighbor := range network.nodeMap[c1].Neighbors {
		if neighbor.Node.Value.(string) == c2 {
			return neighbor.Length, nil
		}
	}
	return -1, errors.New("Invalid Flight Path")
}

// Compute time to go a distance in airplane
func (network Network) computeDistTime(dist int) float32 {
	time := float32(0)
	acceleration := float32(1406.25)

	if dist < 400 { // Compute uniform accel and decel
		half := float32(dist) / float32(2)

		t := math.Sqrt(float64(2 * half / acceleration))

		time += float32(2 * t)

	} else {
		time += 16.0 / 15.0 // Time to accel and decel
		dist -= 400

		time += float32(dist) / float32(750)
	}

	return time
}
