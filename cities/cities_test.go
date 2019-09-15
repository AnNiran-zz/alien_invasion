package cities_test

import (
	"alien_invasion/cities"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewXWorldMap(t *testing.T) {
	newMap := *cities.NewXWorldMap()
	expectedMap := &cities.XWorldMap{
		Existing:  make(map[string]map[string]string),
		Destroyed: make(map[string][]string),
		Aliens:    make(map[string][]string),
	}

	assert.Equal(t, *expectedMap, newMap)
	assert.IsType(t, &newMap, expectedMap)
}

func TestAddCity(t *testing.T) {
	directions := map[string]string{
		"north": "City2",
		"south": "City3",
		"east":  "City4",
	}

	newMap := cities.NewXWorldMap()
	err := newMap.AddCity("City1", directions)

	var errNil error
	assert.Equal(t, errNil, err)

	expectedNewMap := &cities.XWorldMap{
		Existing:  map[string]map[string]string{"City1": {"north": "City2", "south": "City3", "east": "City4"}},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	assert.Equal(t, expectedNewMap, newMap)
	assert.IsType(t, expectedNewMap, newMap)
}

func TestAddCityMissingCityNameError(t *testing.T) {
	directionsCorrect := map[string]string{
		"north": "City2",
		"south": "City3",
		"east":  "City4",
	}

	newMap := cities.NewXWorldMap()
	err := newMap.AddCity("", directionsCorrect)
	expectedErrNoCity := fmt.Errorf("No city name provided")

	assert.Equal(t, expectedErrNoCity, err)
	assert.IsType(t, expectedErrNoCity, err)
}

func TestAddCityIncorrectDirectionError(t *testing.T) {
	directionsWrong := map[string]string{
		"north": "City2",
		"efgh":  "City3",
	}

	newMap := cities.NewXWorldMap()
	err := newMap.AddCity("City1", directionsWrong)
	expectedErrWrongDirection := fmt.Errorf("Incorrect direction provided: efgh")

	assert.Equal(t, expectedErrWrongDirection, err)
	assert.IsType(t, expectedErrWrongDirection, err)
}

func TestAddCityMissingDirectionsError(t *testing.T) {
	directionsMissing := map[string]string{}

	newMap := cities.NewXWorldMap()
	err := newMap.AddCity("City1", directionsMissing)
	expectedErrMissingDirections := fmt.Errorf("No directions provided")

	assert.Equal(t, expectedErrMissingDirections, err)
	assert.Equal(t, expectedErrMissingDirections, err)
}

func TestRemovecity(t *testing.T) {
	xwmapInitial := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3"},
			"City2": {"south": "City1", "east": "City5"},
			"City3": {"north": "City1", "west": "City5"},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15", "alien17"},
		},
	}

	xwmapUpdate, err := xwmapInitial.RemoveCity("City1")

	expectedxwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City2": {"east": "City5"},
			"City3": {"west": "City5"},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15", "alien17"},
		},
	}
	var expectedErr error

	assert.Equal(t, expectedxwmap, xwmapUpdate)
	assert.IsType(t, expectedxwmap, xwmapUpdate)
	assert.Equal(t, expectedErr, err)
	assert.IsType(t, expectedErr, err)
}

func TestRemovecityNonExisting(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {"north": "City1", "west": "City5"},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15", "alien17"},
		},
	}

	xwmapUpdate, err := xwmap.RemoveCity("City7")
	expectedErrNonExistingCity := fmt.Errorf("No city with name City7 exists in the current records")

	assert.Equal(t, expectedErrNonExistingCity, err)
	assert.IsType(t, expectedErrNonExistingCity, err)

	assert.Equal(t, xwmap, xwmapUpdate)
}

func TestExistingCities(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {"north": "City1", "west": "City5"},
		},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	cities := xwmap.ExistingCities()
	sort.Strings(cities)
	expectedCities := []string{"City1", "City2", "City3"}

	assert.Equal(t, expectedCities, cities)
	assert.IsType(t, expectedCities, cities)
}

func TestExistingCitiesEmpty(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing:  map[string]map[string]string{},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	cities := xwmap.ExistingCities()
	var expectedCities []string

	assert.Equal(t, expectedCities, cities)
	assert.IsType(t, expectedCities, cities)
}

func TestDestroyedCitiesEmpty(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing:  map[string]map[string]string{},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	cities := xwmap.DestroyedCities()
	var expectedCities []string

	assert.Equal(t, expectedCities, cities)
	assert.IsType(t, expectedCities, cities)
}

func TestCityRoutes(t *testing.T) {
	newMap := &cities.XWorldMap{
		Existing:  map[string]map[string]string{"City1": {"north": "City2", "south": "City3", "east": "City4"}},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	cityRoutes, err := newMap.CityRoutes("City1")

	var expectedErr error
	expectedCityRoutes := map[string]string{
		"north": "City2",
		"south": "City3",
		"east":  "City4",
	}

	assert.Equal(t, expectedErr, err)
	assert.IsType(t, expectedErr, err)
	assert.Equal(t, expectedCityRoutes, cityRoutes)
	assert.IsType(t, expectedCityRoutes, cityRoutes)
}

func TestCityRoutesMissingCity(t *testing.T) {
	newMap := &cities.XWorldMap{
		Existing:  map[string]map[string]string{"City1": {"north": "City2", "south": "City3", "east": "City4"}},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	cityRoutes, err := newMap.CityRoutes("City9")

	expectedErrMissingCity := fmt.Errorf("No city with name City9 exists in XWorldMap")
	var expectedCityRoutes map[string]string

	assert.Equal(t, expectedErrMissingCity, err)
	assert.IsType(t, expectedErrMissingCity, err)
	assert.Equal(t, expectedCityRoutes, cityRoutes)
	assert.IsType(t, expectedCityRoutes, cityRoutes)
}

func TestAccountAlienLeave(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {"north": "City1", "west": "City5"},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15", "alien17"},
		},
	}

	err := xwmap.AccountAlienLeave("City1", "alien3")
	var expectedErr error

	assert.Equal(t, expectedErr, err)
	assert.IsType(t, expectedErr, err)
}

func TestAccountAlienLeaveNonExistingAlien(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"}},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15", "alien17"},
		},
	}

	err := xwmap.AccountAlienLeave("City1", "alien9")
	expectedErr := fmt.Errorf("alien9 does not exist in any of the XWorld cities")

	assert.Equal(t, expectedErr, err)
	assert.IsType(t, expectedErr, err)
}

func TestAccountAlienLeaveNonExistingCity(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"}},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15", "alien17"},
		},
	}

	err := xwmap.AccountAlienLeave("City90", "alien9")
	expectedErr := fmt.Errorf("City90 is not among existing cities in XWorld keeping aliens")

	assert.Equal(t, expectedErr, err)
	assert.IsType(t, expectedErr, err)
}
