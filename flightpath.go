package main

// Type to unmarshal flightpath data
type FlightPath struct {
	Ports    []string `json:"ports"`
	Distance int      `json:"distance"`
}
