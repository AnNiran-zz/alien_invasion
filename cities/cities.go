package cities

import (
	"fmt"
)

// XWorldMap contains data about each of existing cities and its existing outisde directions
// format is: direction=destination
// Cities keys and their direction references are updated continously along with to aliens moves
type XWorldMap struct {
	Existing  map[string]map[string]string
	Destroyed map[string][]string
	Aliens    map[string][]string
}

// XWorld is an instance of XWorldMap and its pointer is used across
// the program to update map data
var XWorld *XWorldMap

// NewXWorldMap creates a map for all cities and corresponding data received from input
func NewXWorldMap() *XWorldMap {
	return &XWorldMap{
		Existing:  make(map[string]map[string]string),
		Destroyed: make(map[string][]string),
		Aliens:    make(map[string][]string),
	}
}

// AddCity adds a new city to the XWorld map
// Includes checks for insufficient city direction references and empty city name
// Directions is a map containing city names in each direction and is expected to
// contain empty directions only for existing outroutes
func (xwm *XWorldMap) AddCity(cityName string, directions map[string]string) error {

	// Check if city name is provided
	if cityName == "" {
		return fmt.Errorf("No city name provided")
	}

	// Check if directions are correnctly provided
	if len(directions) < 1 {
		return fmt.Errorf("No directions provided")
	}

	for dir := range directions {
		if dir != "north" && dir != "south" && dir != "west" && dir != "east" {
			return fmt.Errorf("Incorrect direction provided: %s", dir)
		}
	}

	newCity := make(map[string]string)

	if directions["east"] != "" {
		newCity["east"] = directions["east"]
	}

	if directions["west"] != "" {
		newCity["west"] = directions["west"]
	}

	if directions["north"] != "" {
		newCity["north"] = directions["north"]
	}

	if directions["south"] != "" {
		newCity["south"] = directions["south"]
	}
	xwm.Existing[cityName] = newCity

	return nil
}

// RemoveCity removes a city and its direction references form the map
func (xwm *XWorldMap) RemoveCity(cityName string) (*XWorldMap, error) {
	if _, ok := xwm.Existing[cityName]; !ok {
		return xwm, fmt.Errorf("No city with name %s exists in the current records", cityName)
	}

	// Remove routes going to the city
	// Obtain city existing map and keep it in a separate address
	cityMap := xwm.Existing[cityName]
	cityRefMap := make(map[string]string)
	for _, cityDest := range cityMap {
		// Locate each city that is pointing to the one we need to destroy
		cityRefMap = xwm.Existing[cityDest]
		for route, destToRemove := range cityRefMap {
			// Remove the route that points to the city to be destroyed
			if destToRemove == cityName {
				delete(cityRefMap, route)
				fmt.Printf("Reference to %s has been removed from %s \n", cityName, cityDest)
				break
			}
		}
		// Update xwm map with record without the route
		xwm.Existing[cityDest] = cityRefMap
	}
	delete(xwm.Existing, cityName)

	fmt.Printf("%s has been removed from XWorld map.\n", cityName)
	return xwm, nil
}

// ExistingCities return a slice of existing city names
func (xwm *XWorldMap) ExistingCities() []string {
	var cities []string
	for name := range xwm.Existing {
		cities = append(cities, name)
	}
	return cities
}

// DestroyedCities return a slice of destroyed cities
func (xwm *XWorldMap) DestroyedCities() []string {
	var cities []string
	for name := range xwm.Destroyed {
		cities = append(cities, name)
	}
	return cities
}

// CityRoutes returns all routes from an existing city
func (xwm *XWorldMap) CityRoutes(cityName string) (map[string]string, error) {
	if _, ok := xwm.Existing[cityName]; !ok {
		return nil, fmt.Errorf("No city with name %s exists in XWorldMap", cityName)
	}

	return xwm.Existing[cityName], nil
}

// AccountAlienLeave takes into account an alien has left the city
func (xwm *XWorldMap) AccountAlienLeave(cityName, alienName string) error {
	present := xwm.Aliens[cityName]
	if _, ok := xwm.Aliens[cityName]; !ok {
		return fmt.Errorf("%s is not among existing cities in XWorld keeping aliens", cityName)
	}

	for value, name := range present {
		if name == alienName {
			copy(present[value:], present[value+1:])
			present[len(present)-1] = ""
			present = present[:len(present)-1]

			xwm.Aliens[cityName] = present
			return nil
		}
	}
	return fmt.Errorf("%s does not exist in any of the XWorld cities", alienName)
}
