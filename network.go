package main

import (
	"encoding/json"
	graphlib "github.com/ehaydenr/algorithms/graph"
	"io/ioutil"
)

type Network struct {
	Airports    map[string]Airport
	FlightPaths []FlightPath
	graph       *graphlib.Graph
	nodeMap     map[string]*graphlib.Vertex
}

// Constructor
func NewNetwork() *Network {
	data, err := ioutil.ReadFile("res/map_data.json")
	if err != nil {
		panic(err)
	}

	config := Configuration{}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	net := &Network{}

	airportMap := make(map[string]Airport)
	for _, airport := range config.Airports {
		airportMap[airport.Code] = airport
	}

	net.Airports = airportMap
	net.FlightPaths = config.FlightPaths
	net.buildGraph()

	return net
}

// Merge JSON
func (network *Network) MergeJSON(config Configuration) {
	for _, airport := range config.Airports {
		network.Airports[airport.Code] = airport
	}
	for _, path := range config.FlightPaths {
		network.FlightPaths = append(network.FlightPaths, path)
	}
	network.buildGraph()
}

// Generate resources from Airport and Flightpath Data
func (network *Network) buildGraph() {
	graph := make(graphlib.Graph, len(network.Airports))
	nodeMap := make(map[string]*graphlib.Vertex)

	i := 0
	for code := range network.Airports {
		graph[i] = &graphlib.Vertex{
			Id:        i,
			Value:     code,
			Neighbors: make([]graphlib.Neighbor, 0),
		}

		nodeMap[code] = graph[i]
		i++
	}

	for _, path := range network.FlightPaths {
		v1 := nodeMap[path.Ports[0]]
		v2 := nodeMap[path.Ports[1]]
		v1.Neighbors = append(v1.Neighbors, graphlib.Neighbor{v2, path.Distance})
		v2.Neighbors = append(v2.Neighbors, graphlib.Neighbor{v1, path.Distance})
	}

	network.graph = &graph
	network.nodeMap = nodeMap
}

// Remove City
func (network *Network) removeCity(code string) {
	// Remove city from airport map
	delete(network.Airports, code)

	// Remove Flight Paths
	newPaths := make([]FlightPath, 0, len(network.FlightPaths))
	for _, path := range network.FlightPaths {
		if path.Ports[0] != code || path.Ports[1] != code {
			newPaths = append(newPaths, path)
		}
	}
	network.FlightPaths = newPaths

	// Could be more optimized by not deleting everything, but let's just rebuild for now.
	network.buildGraph()
}

// Add City
func (network *Network) addCity(airport Airport) {
	network.Airports[airport.Code] = airport
	network.buildGraph()
}

// Update City
func (network *Network) updateCity(airport Airport) {
	network.Airports[airport.Code] = airport
	network.buildGraph()
}

// Remove Route
func (network *Network) removeRoute(port1, port2 string) {
	newFlightPaths := make([]FlightPath, 0, len(network.FlightPaths))
	for _, fp := range network.FlightPaths {
		if fp.Ports[0] != port1 && fp.Ports[1] != port2 {
			newFlightPaths = append(newFlightPaths, fp)
		}
	}
	network.FlightPaths = newFlightPaths
	network.buildGraph()
}

// Add Route
func (network *Network) addRoute(fp FlightPath) {
	network.FlightPaths = append(network.FlightPaths, fp)
	network.buildGraph()
}
