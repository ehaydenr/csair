package main

// Type to unmarshal coordinates
type Coordinates struct {
	S int `json:"S"`
	W int `json:"W"`
}

// Type to unmarshal airport data
type Airport struct {
	Code       string      `json:"code"`
	Name       string      `json:"name"`
	Country    string      `json:"country"`
	Continent  string      `json:"continent"`
	Timezone   float32     `json:"timezone"`
	Coords     Coordinates `json:"coordinates"`
	Population int         `json:"population"`
	Region     int         `json:"region"`
}

// Thrown if querying nonexistent arirport
type NoSuchAirportError string

// Implement Error interface
func (e NoSuchAirportError) Error() string {
	return string(e)
}
