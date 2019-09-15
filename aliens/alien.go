package aliens

import (
	"alien_invasion/cities"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Alien struct keeps record of an existing alien
// holds data that is going to be used during the invasion and status updates
type alien struct {
	name      string
	position  string
	moveCount int
	status    string
}

// AliensRelease keeps track of all released aliens over XWorld
//
// Alien status "destroyed" prevails over statuses: "inoperation", "trapped" and "exausted"
// Status "trapped" and "exausted" can be set only to an alien that has a status "inoperation"
// Once an alien status is changed from "inoperation", it is never returned to that
//
// inOperation a record of all currently existing aliens, which are able to to take a consequent
// move across the map - are not trapped, destroyed or exausted their moves
//
// trapped keeps a record of all aliens that are trapped in a city in XWorld map
//
// exausted keeps track of all aliens that have exausted their allowed number of moves
//
// destroyed keeps a record about all aliens that have been destroyed
type aliensRelease struct {
	all         []*alien
	inOperation []*alien
	trapped     []*alien
	exausted    []*alien
	destroyed   []*alien
}

// Defines maximum number of moves an alien can make
// Defines maximum number of moves an alien can make
const (
	maxMovementsAllowed = 10000

	alienInOperationStatus = "inoperation"
	alienTrappedStatus     = "trapped"
	alienExaustedStatus    = "exausted"
	alienDestroyedStatus   = "destroyed"
)

// aliensReleased hold cache data for all aliens statuses
var aliensReleased = &aliensRelease{}

// calculateInitialPosition returns initial random position of an alien on the XWorld map
func calculateInitialPosition(xwm *cities.XWorldMap) string {
	potPositions := xwm.ExistingCities()
	rand.Seed(time.Now().Unix())
	randPosition := rand.Int() % len(potPositions)

	return potPositions[randPosition]
}

// newAlien creates a new alien using its number to record a name for it
// and calculates it initial position
func newAlien(alienN int, xwm *cities.XWorldMap) *alien {
	landAt := calculateInitialPosition(xwm)
	alienName := "alien" + strconv.Itoa(alienN)

	return &alien{
		name:      alienName,
		position:  landAt,
		moveCount: 0,
		status:    alienInOperationStatus,
	}
}

// sendToNextPosition calculates next random position of an alien
// on the XWorld map, depending on the current records of existing cities,
// and moves the alien to it, taking into account to update records for the move
func (alien *alien) sendToNextPosition(xwm *cities.XWorldMap) (*cities.XWorldMap, error) {
	// check if alien has exausted all possible moves count
	if alien.moveCount == maxMovementsAllowed {
		alien.accountExausted()
		return nil, fmt.Errorf("%s cannot be moved because has exausted allowed moves count", alien.name)
	}

	// caclulate next direction by randimizing current options
	var nextDir string
	var randDirections []string
	for dir := range xwm.Existing[alien.position] {
		randDirections = append(randDirections, dir)
	}

	if len(randDirections) == 0 {
		// if there are no further direction - the alien is trapped
		if !isInTrappedRecords(alien) {
			alien.accountTrapped()
		}
		return nil, fmt.Errorf("%s cannot be moved because it is trapped", alien.name)

	} else if len(randDirections) == 1 {
		// if there is one direction left - choose it
		nextDir = randDirections[0]

	} else {
		// in all other cases - choose random position
		rand.Seed(time.Now().Unix())
		nextPos := rand.Int() % len(randDirections)

		nextDir = randDirections[nextPos]
	}

	// Account alien left the city in the map
	err := xwm.AccountAlienLeave(alien.position, alien.name)
	if err != nil {
		return nil, err
	}

	alien.position = xwm.Existing[alien.position][nextDir]
	alien.moveCount++

	xwm.Aliens[alien.position] = append(xwm.Aliens[alien.position], alien.name)
	fmt.Printf("%s is moving to %s \n", alien.name, alien.position)

	return xwm, nil
}

// isTrapped calculates the current status of an alien - if there are
// any possible routes out of the current position or not
func (alien *alien) isTrapped(xwm *cities.XWorldMap) bool {
	// Get status of current position - check if entity is trapped
	curRoutes := xwm.Existing[alien.position]

	if len(curRoutes) < 1 {
		alien.status = alienTrappedStatus
		return true
	}
	return false
}

// isInTrappedRecords checks if an alien is recordred as trapped
func isInTrappedRecords(alien *alien) bool {
	for _, alienTrapped := range aliensReleased.trapped {
		if alienTrapped.name == alien.name {
			return true
		}
	}
	return false
}

// accountTrapped remove alien from the list of alien in operation,
// which will terminate all future possible movements for it
func (alien *alien) accountTrapped() {
	inOperationEntities := aliensReleased.inOperation

	for value, operational := range inOperationEntities {
		if alien.name == operational.name {
			alien.status = alienTrappedStatus
			aliensReleased.trapped = append(aliensReleased.trapped, alien)

			copy(inOperationEntities[value:], inOperationEntities[value+1:])
			inOperationEntities[len(inOperationEntities)-1] = nil
			inOperationEntities = inOperationEntities[:len(inOperationEntities)-1]

			aliensReleased.inOperation = inOperationEntities
			break
		}
	}

	for _, released := range aliensReleased.all {
		if released.name == alien.name {
			released.status = alienTrappedStatus
			break
		}
	}
}

// accountExausted remove alien from operational, which will terminate all
// future possible movements for it
func (alien *alien) accountExausted() {
	inOperationEntities := aliensReleased.inOperation

	for value, operational := range inOperationEntities {
		if alien.name == operational.name {
			alien.status = alienExaustedStatus
			aliensReleased.exausted = append(aliensReleased.exausted, alien)

			copy(inOperationEntities[value:], inOperationEntities[value+1:])
			inOperationEntities[len(inOperationEntities)-1] = nil
			inOperationEntities = inOperationEntities[:len(inOperationEntities)-1]

			aliensReleased.inOperation = inOperationEntities
			break
		}
	}

	for _, released := range aliensReleased.all {
		if released.name == alien.name {
			released.status = alienExaustedStatus
			break
		}
	}
}

// accountDestroyed tries to destroy an alien
// and uses separate functions depending on its status
func accountDestroyed(name string) bool {
	var success bool
	for _, alive := range aliensReleased.all {
		if name == alive.name {
			switch alive.status {
			case alienInOperationStatus:
				success = destroyInOperationAlien(alive)

			case alienTrappedStatus:
				success = destroyTrappedAlien(alive)

			case alienExaustedStatus:
				success = destroyExaustedAlien(alive)

			default:
				//
			}

			if !success {
				return false
			}

			for _, released := range aliensReleased.all {
				if released.name == name {
					released.status = alienDestroyedStatus
				}
				break
			}
		}
	}
	fmt.Printf("%s has been destroyed.\n", name)
	return true
}

// destroyTrappedAlien removes alien from slice containing currently trapped aliens
func destroyTrappedAlien(alien *alien) bool {
	trappedEntities := aliensReleased.trapped

	for value, trapped := range trappedEntities {
		if trapped.name == alien.name {
			alien.status = alienDestroyedStatus
			aliensReleased.destroyed = append(aliensReleased.destroyed, alien)

			copy(trappedEntities[value:], trappedEntities[value+1:])
			trappedEntities[len(trappedEntities)-1] = nil
			trappedEntities = trappedEntities[:len(trappedEntities)-1]

			aliensReleased.trapped = trappedEntities
			return true
		}
	}
	return false
}

// destroyExaustedAlien remove alien from slice containing currently exausted aliens
func destroyExaustedAlien(alien *alien) bool {
	exaustedEntities := aliensReleased.exausted

	for value, exausted := range exaustedEntities {
		if alien.name == exausted.name {
			alien.status = alienDestroyedStatus
			aliensReleased.destroyed = append(aliensReleased.destroyed, alien)

			copy(exaustedEntities[value:], exaustedEntities[value+1:])
			exaustedEntities[len(exaustedEntities)-1] = nil
			exaustedEntities = exaustedEntities[:len(exaustedEntities)-1]

			aliensReleased.exausted = exaustedEntities
			return true
		}
	}

	return false
}

// destroyInOperationAlien remove alien from slice containing currently inOperation aliens
func destroyInOperationAlien(alien *alien) bool {
	inOperationEntities := aliensReleased.inOperation

	for value, inOper := range inOperationEntities {
		if alien.name == inOper.name {
			alien.status = alienDestroyedStatus
			aliensReleased.destroyed = append(aliensReleased.destroyed, alien)

			copy(inOperationEntities[value:], inOperationEntities[value+1:])
			inOperationEntities[len(inOperationEntities)-1] = nil
			inOperationEntities = inOperationEntities[:len(inOperationEntities)-1]

			aliensReleased.inOperation = inOperationEntities
			return true
		}
	}

	return false
}
