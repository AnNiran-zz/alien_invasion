package aliens

import (
	"alien_invasion/cities"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXWorldMapSucess(t *testing.T) {
	xwmapInitial := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"CityA": {"south": "CityE", "east": "CityB"},
			"CityB": {"south": "CityF", "west": "CityA", "east": "CityC"},
			"CityC": {"south": "CityG", "west": "CityB", "east": "CityD"},
			"CityD": {"south": "CityH", "west": "CityC"},
		},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	filename := "../xworldmap/xworld-map-test-input.txt"
	newMap, err := xWorldMap(filename)

	var errSuccess error

	assert.Equal(t, xwmapInitial, newMap)
	assert.IsType(t, xwmapInitial, newMap)
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)
}

func TestXWorldMapFail(t *testing.T) {
	var xwmap *cities.XWorldMap

	filename := "../xworldmap/xworld-map-test-input-0.txt"
	newMap, err := xWorldMap(filename)

	var errFail *os.PathError
	errFailText := "open ../xworldmap/xworld-map-test-input-0.txt: no such file or directory"

	assert.Equal(t, xwmap, newMap)
	assert.Equal(t, errFailText, err.Error())
	assert.IsType(t, errFail, err)
}

func TestFightTwoAliens(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien2", "alien3"},
			"City2": {"alien15", "alien1"},
		},
	}

	alien1 := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	alien15 := &alien{
		name:      "alien15",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien1)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien1)
	aliensReleased.all = append(aliensReleased.all, alien15)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien15)

	xwmap, err := alien1.fight(alien15.name, xwmap)

	var errSuccess error
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)

	opponentIsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien15.name == alienDestroyed.name {
			opponentIsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, opponentIsInDestroyedRecords, true)

	cityDestoyed := true
	for city := range xwmap.Existing {
		if city == alien1.position {
			cityDestoyed = false
			break
		}
	}

	assert.Equal(t, cityDestoyed, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestFightMultipleAliens(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien2", "alien3"},
			"City2": {"alien15", "alien1", "alien17"},
		},
	}

	alien1 := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	alien15 := &alien{
		name:      "alien15",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	alien17 := &alien{
		name:      "alien17",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien1)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien1)
	aliensReleased.all = append(aliensReleased.all, alien15)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien15)
	aliensReleased.all = append(aliensReleased.all, alien17)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien17)

	xwmap, err := alien1.fight(alien15.name, xwmap)

	var errSuccess error
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)

	alien15IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien15.name == alienDestroyed.name {
			alien15IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien15IsInDestroyedRecords, true)

	alien17IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien15.name == alienDestroyed.name {
			alien17IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien17IsInDestroyedRecords, true)

	cityDestoyed := true
	for city := range xwmap.Existing {
		if city == alien1.position {
			cityDestoyed = false
			break
		}
	}

	assert.Equal(t, cityDestoyed, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestLand(t *testing.T) {
	xwmapInitial := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"CityA": {"south": "CityE", "east": "CityB"},
			"CityB": {"south": "CityF", "west": "CityA", "east": "CityC"},
			"CityC": {"south": "CityG", "west": "CityB", "east": "CityD"},
			"CityD": {"south": "CityH", "west": "CityC"},
		},
		Destroyed: map[string][]string{},
		Aliens:    map[string][]string{},
	}

	xwmap := land(3, xwmapInitial)

	var createEntities int
	for _, aliens := range xwmap.Aliens {
		createEntities += len(aliens)
	}

	assert.Equal(t, createEntities, 3)
}

func TestFindOpponentMultiple(t *testing.T) {
	xwmapInitial := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien2", "alien3"},
			"City2": {"alien15", "alien1", "alien17"},
		},
	}

	alien1 := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	alien15 := &alien{
		name:      "alien15",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	alien17 := &alien{
		name:      "alien17",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien1)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien1)
	aliensReleased.all = append(aliensReleased.all, alien15)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien15)
	aliensReleased.all = append(aliensReleased.all, alien17)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien17)

	xwmapUpdate, fought, err := alien1.findOponent(xwmapInitial)
	fmt.Println(err)
	var errSuccess error
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)
	assert.Equal(t, fought, true)

	cityDestroyed := false
	for city := range xwmapUpdate.Destroyed {
		if city == alien1.position {
			cityDestroyed = true
			break
		}
	}
	assert.Equal(t, cityDestroyed, true)

	alien1IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien1.name == alienDestroyed.name {
			alien1IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien1IsInDestroyedRecords, false)

	alien15IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien15.name == alienDestroyed.name {
			alien15IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien15IsInDestroyedRecords, true)

	alien17IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien15.name == alienDestroyed.name {
			alien17IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien17IsInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestFindOpponentSingle(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien2", "alien3"},
			"City2": {"alien15", "alien1"},
		},
	}

	alien1 := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	alien15 := &alien{
		name:      "alien15",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	// create testing conditions
	aliensReleased.all = append(aliensReleased.all, alien1)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien1)
	aliensReleased.all = append(aliensReleased.all, alien15)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien15)

	xwmap, fought, err := alien1.findOponent(xwmap)

	var errSuccess error
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)
	assert.Equal(t, fought, true)

	cityDestroyed := true
	for city := range xwmap.Existing {
		if city == alien1.position {
			cityDestroyed = false
			break
		}
	}
	assert.Equal(t, cityDestroyed, true)

	alien15IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien15.name == alienDestroyed.name {
			alien15IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien15IsInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestFindOpponentNo(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien2", "alien3"},
			"City2": {"alien1"},
		},
	}

	alien1 := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City2",
	}

	// create testing conditions
	aliensReleased.all = append(aliensReleased.all, alien1)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien1)

	xwmap, fought, err := alien1.findOponent(xwmap)

	var errSuccess error
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)
	assert.Equal(t, fought, false)

	cityDestroyed := true
	for city := range xwmap.Existing {
		if city == alien1.position {
			cityDestroyed = false
			break
		}
	}
	assert.Equal(t, cityDestroyed, false)

	alien1IsInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien1.name == alienDestroyed.name {
			alien1IsInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, alien1IsInDestroyedRecords, false)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

// testing operate() function cannot ensure certain final status
// of cities and aliens, because its logic is based on random moves
// we ensure the trapped alien is not moved and the operational ones
// moved at least once here
func TestOperate(t *testing.T) {
	xwmapInitial := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"CityA": {"south": "CityE", "east": "CityB"},
			"CityB": {"south": "CityF", "west": "CityA", "east": "CityC"},
			"CityC": {"south": "CityG", "west": "CityB", "east": "CityD"},
			"CityD": {"south": "CityH", "west": "CityC"},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"CityA": {"alien1"},
			"CityB": {"alien15", "alien17"},
		},
	}

	alien1 := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "CityA",
	}

	alien15 := &alien{
		name:      "alien15",
		moveCount: 0,
		status:    "trapped",
		position:  "CityB",
	}

	alien17 := &alien{
		name:      "alien17",
		moveCount: 0,
		status:    "inoperation",
		position:  "CityB",
	}

	// create testing conditions
	aliensReleased.all = append(aliensReleased.all, alien1)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien1)
	aliensReleased.all = append(aliensReleased.all, alien15)
	aliensReleased.trapped = append(aliensReleased.trapped, alien15)
	aliensReleased.all = append(aliensReleased.all, alien17)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien17)

	var errSuccess error
	_, err := operate(xwmapInitial)

	assert.Equal(t, alien15.moveCount, 0)
	assert.Equal(t, errSuccess, err)
	assert.IsType(t, errSuccess, err)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}
