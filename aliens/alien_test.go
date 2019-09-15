package aliens

import (
	"alien_invasion/cities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlien(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15"},
		},
	}

	newAlien := newAlien(30, xwmap)

	expectedName := "alien30"
	expectedStatus := "inoperation"
	expectedmoveCount := 0

	positionSet := false
	cities := xwmap.ExistingCities()
	for _, city := range cities {
		if newAlien.position == city {
			positionSet = true
			break
		}
	}

	assert.Equal(t, expectedName, newAlien.name)
	assert.IsType(t, expectedName, newAlien.name)
	assert.Equal(t, expectedStatus, newAlien.status)
	assert.Equal(t, expectedmoveCount, newAlien.moveCount)
	assert.Equal(t, positionSet, true)
}

func TestIsTrappedTrue(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15"},
		},
	}

	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City3",
	}

	resultTrapped := alien.isTrapped(xwmap)
	expectedResultTrapped := true

	assert.Equal(t, resultTrapped, expectedResultTrapped)
	assert.IsType(t, resultTrapped, expectedResultTrapped)
}

func TestIsTrappedFalse(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15"},
		},
	}

	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City1",
	}

	resultTrapped := alien.isTrapped(xwmap)
	expectedResultTrapped := false

	assert.Equal(t, resultTrapped, expectedResultTrapped)
	assert.IsType(t, resultTrapped, expectedResultTrapped)
}

func TestAccountTrapped(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	alien.accountTrapped()

	expectedAlienStatus := "trapped"
	inTrappedList := false
	for _, alienTrapped := range aliensReleased.trapped {
		if alien.name == alienTrapped.name {
			inTrappedList = true
			break
		}
	}

	assert.Equal(t, expectedAlienStatus, alien.status)
	assert.Equal(t, inTrappedList, true)

	inInOperationList := false
	for _, alienInOperation := range aliensReleased.inOperation {
		if alien.name == alienInOperation.name {
			inInOperationList = true
			break
		}
	}

	assert.Equal(t, inInOperationList, false)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestDestroyTrappedAlienSuccess(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "trapped",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.trapped = append(aliensReleased.trapped, alien)
	destroyed := destroyTrappedAlien(alien)

	expectedStatus := "destroyed"
	expectedResult := true

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	isInTrappedRecords := isInTrappedRecords(alien)
	assert.Equal(t, false, isInTrappedRecords)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestDestroyTrappedAlienFail(t *testing.T) {
	alien := &alien{
		name:      "alien0",
		moveCount: 0,
		status:    "trapped",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.trapped = append(aliensReleased.trapped, alien)
	destroyed := destroyTrappedAlien(alien)

	expectedStatus := "destroyed"
	expectedResult := true

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	isInTrappedRecords := isInTrappedRecords(alien)
	assert.Equal(t, false, isInTrappedRecords)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestAccountExausted(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	alien.accountExausted()

	expectedStatus := "exausted"
	inExaustedList := false

	for _, alienExausted := range aliensReleased.exausted {
		if alien.name == alienExausted.name {
			inExaustedList = true
			break
		}
	}

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, inExaustedList, true)

	inInOperationList := false
	for _, alienInOperation := range aliensReleased.inOperation {
		if alien.name == alienInOperation.name {
			inInOperationList = true
			break
		}
	}

	assert.Equal(t, inInOperationList, false)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestDestroyExaustedAlien(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "exausted",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.exausted = append(aliensReleased.exausted, alien)
	destroyed := destroyExaustedAlien(alien)

	expectedStatus := "destroyed"
	expectedResult := true

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestDestroyInOperationAlien(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	destroyed := destroyInOperationAlien(alien)

	expectedStatus := "destroyed"
	expectedResult := true

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestAccountDestroyedInOperationEntity(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City1",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	destroyed := accountDestroyed(alien.name)

	expectedStatus := "destroyed"
	expectedResult := true

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestAccountDestroyedTrappedEntity(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "trapped",
		position:  "City1",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.trapped = append(aliensReleased.trapped, alien)
	destroyed := accountDestroyed(alien.name)

	//expectedStatus := "destroyed"
	expectedResult := true

	//assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestAccountDestroyedExaustedEntity(t *testing.T) {
	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "exausted",
		position:  "City1",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.exausted = append(aliensReleased.exausted, alien)
	destroyed := accountDestroyed(alien.name)

	expectedStatus := "destroyed"
	expectedResult := true

	assert.Equal(t, expectedStatus, alien.status)
	assert.Equal(t, expectedResult, destroyed)
	assert.IsType(t, expectedResult, destroyed)

	isInDestroyedRecords := false
	for _, alienDestroyed := range aliensReleased.destroyed {
		if alien.name == alienDestroyed.name {
			isInDestroyedRecords = true
			break
		}
	}
	assert.Equal(t, isInDestroyedRecords, true)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestSendToNextPositionInOperationEntity(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15"},
		},
	}

	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City1",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	xwmap.Aliens[alien.position] = append(xwmap.Aliens[alien.position], alien.name)

	alienInitialPosition := alien.position
	_, result := alien.sendToNextPosition(xwmap)

	var errSuccess error
	assert.Equal(t, result, nil)
	assert.IsType(t, result, errSuccess)

	moved := false
	if alien.position != alienInitialPosition {
		moved = true
	}

	assert.Equal(t, moved, true)

	arrivedInNewCity := false
	for _, alienPresent := range xwmap.Aliens[alien.position] {
		if alienPresent == alien.name {
			arrivedInNewCity = true
		}
	}

	assert.Equal(t, arrivedInNewCity, true)
	assert.Equal(t, alien.moveCount, 1)

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestSendToNextPositionTrappedEntity(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15"},
		},
	}

	alien := &alien{
		name:      "alien1",
		moveCount: 0,
		status:    "inoperation",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	xwmap.Aliens[alien.position] = append(xwmap.Aliens[alien.position], alien.name)

	alienInitialPosition := alien.position
	_, result := alien.sendToNextPosition(xwmap)
	errCannotMove := fmt.Errorf("alien1 cannot be moved because it is trapped")

	assert.Equal(t, result, errCannotMove)
	assert.Equal(t, alienInitialPosition, alien.position)
	assert.Equal(t, alien.moveCount, 0)

	inCityRecords := false
	for _, alienPresent := range xwmap.Aliens[alien.position] {
		if alienPresent == alien.name {
			inCityRecords = true
		}
	}

	assert.Equal(t, inCityRecords, true)
	assert.Equal(t, alien.status, "trapped")

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}

func TestSendToNextPositionExaustedEntity(t *testing.T) {
	xwmap := &cities.XWorldMap{
		Existing: map[string]map[string]string{
			"City1": {"north": "City2", "south": "City3", "east": "City4"},
			"City2": {"south": "City1", "west": "City4", "east": "City5"},
			"City3": {},
		},
		Destroyed: map[string][]string{},
		Aliens: map[string][]string{
			"City1": {"alien1", "alien2", "alien3"},
			"City2": {"alien15"},
		},
	}

	alien := &alien{
		name:      "alien1",
		moveCount: 10000,
		status:    "inoperation",
		position:  "City3",
	}

	// create test conditions
	aliensReleased.all = append(aliensReleased.all, alien)
	aliensReleased.inOperation = append(aliensReleased.inOperation, alien)
	xwmap.Aliens[alien.position] = append(xwmap.Aliens[alien.position], alien.name)

	alienInitialPosition := alien.position
	_, result := alien.sendToNextPosition(xwmap)
	errCannotMove := fmt.Errorf("alien1 cannot be moved because has exausted allowed moves count")

	assert.Equal(t, result, errCannotMove)
	assert.Equal(t, alienInitialPosition, alien.position)

	inCityRecords := false
	for _, alienPresent := range xwmap.Aliens[alien.position] {
		if alienPresent == alien.name {
			inCityRecords = true
		}
	}

	assert.Equal(t, inCityRecords, true)
	assert.Equal(t, alien.status, "exausted")

	// destroy test conditions
	aliensReleased = &aliensRelease{}
}
