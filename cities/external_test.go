package cities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadXWorlCitiesDataSuccess(t *testing.T) {
	xwmapInitial := &XWorldMap{
		Existing:  map[string]map[string]string{},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	filename := "../xworldmap/xworld-map-test-input.txt"
	err := xwmapInitial.ReadXWorldCitiesData(filename)

	var errSuccess error
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)

	xwmapExpected := &XWorldMap{
		Existing: map[string]map[string]string{
			"CityA": {"south": "CityE", "east": "CityB"},
			"CityB": {"south": "CityF", "west": "CityA", "east": "CityC"},
			"CityC": {"south": "CityG", "west": "CityB", "east": "CityD"},
			"CityD": {"south": "CityH", "west": "CityC"},
		},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	assert.Equal(t, xwmapInitial, xwmapExpected)
}

func TestReadXWorlCitiesDataNonExistingFile(t *testing.T) {
	xwmapInitial := &XWorldMap{
		Existing:  map[string]map[string]string{},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	filename := "../xworldmap/xworld-map-test-input-0.txt"
	err := xwmapInitial.ReadXWorldCitiesData(filename)

	errNoExistingFile := "open ../xworldmap/xworld-map-test-input-0.txt: no such file or directory"
	assert.Equal(t, errNoExistingFile, err.Error())

	xwmapExpected := &XWorldMap{
		Existing:  map[string]map[string]string{},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	assert.Equal(t, xwmapInitial, xwmapExpected)
}
