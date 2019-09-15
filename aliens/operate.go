package aliens

import (
	"alien_invasion/cities"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// xWorldMap uses cities.XWorld instance to save data read from input file
// in order to be used for further reading when aliens operate on the cities
func xWorldMap(filename string) (*cities.XWorldMap, error) {
	// assign cities package variable to a new map
	cities.XWorld = cities.NewXWorldMap()

	err := cities.XWorld.ReadXWorldCitiesData(filename)
	if err != nil {
		return nil, err
	}

	// return a pointer to the package variable
	// that will be used during program execution
	return cities.XWorld, nil
}

// fight represent a fight between two aliens in a city, which destroys the city,
// both of them and all other aliens that are currently in it
func (alien *alien) fight(opponent string, xwm *cities.XWorldMap) (*cities.XWorldMap, error) {
	fmt.Printf("%s and %s are fighting in %s ... \n", alien.name, opponent, alien.position)

	// Account all aliens in the city as destroyed
	presentAliens := xwm.Aliens[alien.position]

	for _, presentAlien := range presentAliens {
		_ = accountDestroyed(presentAlien)
	}

	// Destroy city and add it to record of destroyed
	existingCities := xwm.ExistingCities()
	for _, cityName := range existingCities {
		if cityName == alien.position {
			xwm.Destroyed[alien.position] = []string{alien.name, opponent}
			break
		}
	}
	delete(xwm.Aliens, alien.position)

	// Remove city and all directions pointing to it from the
	xwmUpdate, err := xwm.RemoveCity(alien.position)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%s has been destroyed by %s and %s!\n", alien.position, alien.name, opponent)
	return xwmUpdate, nil
}

// land releases the aliens on XWorld cities
//
// Creates requested number of aliens and saves them to the cache record
// for aliens in operation
// Updates cities.XWorld.Aliens data with the initial posirions
func land(aliens int, xwm *cities.XWorldMap) *cities.XWorldMap {
	for alienCount := 0; alienCount < aliens; alienCount++ {
		alien := newAlien(alienCount, xwm)
		aliensReleased.all = append(aliensReleased.all, alien)
		aliensReleased.inOperation = append(aliensReleased.inOperation, alien)

		xwm.Aliens[alien.position] = append(xwm.Aliens[alien.position], alien.name)
	}

	fmt.Printf("%d aliens created and landed across XWorld cities ", aliens)
	return xwm
}

// findOpponent tries to find another alien in the current city to fight with
// if aliens are more than one - select a random one and initiates a fight with it
// After the fight - all corresponding record changes are taken into account by the fight() function
// If there are no aliens to fight - move to next position
func (alien *alien) findOponent(xwm *cities.XWorldMap) (*cities.XWorldMap, bool, error) {
	aliensInCity := xwm.Aliens[alien.position]

	// exclude my name from the present aliens in the city to get list of potential oponents
	for value, presentAlien := range aliensInCity {
		if presentAlien == alien.name {
			aliensInCity = append(aliensInCity[:value], aliensInCity[value+1:]...)
			break
		}
	}

	var opponent string
	if len(aliensInCity) > 0 {
		// Choose a random opponent among the rest of aliens in the city
		if len(aliensInCity) > 1 {
			rand.Seed(time.Now().Unix())
			randSelect := rand.Int() % len(aliensInCity)

			opponent = aliensInCity[randSelect]

		} else {
			// If there is one other alien - choose it as an opponent
			opponent = aliensInCity[0]
		}

		// Initiate a fight with the other alien
		xwmUpdate, err := alien.fight(opponent, xwm)
		if err != nil {
			return xwmUpdate, false, err
		}

		return xwmUpdate, true, nil
	}

	return xwm, false, nil
}

// operate starts aliens random movements across the XWorld cities
//
// Each entity movement is represented by a loop iteration with a maximum count of maxMovementsAllowed
// if number of aliens that are able to operate becomes zero before reachin the maxMovementsAllowed
// loops - the loops breaks earlier
func operate(xwm *cities.XWorldMap) (*cities.XWorldMap, error) {

	// Update XWorld map and aliens status
	var err error

	for i := 0; i <= maxMovementsAllowed; i++ {
		// Make a move for each alien
		for _, alien := range aliensReleased.all {
			// if alien is destoryed move to the next loop
			if alien.status == alienDestroyedStatus {
				continue
			}

			// check if alien is trapped after the previous steps
			if alien.isTrapped(xwm) {
				if !isInTrappedRecords(alien) {
					alien.accountTrapped()
				}
			}

			// Check alien status:
			// proceed with moving to next position only of the alien has a status "inoperation"
			if alien.status == alienInOperationStatus {
				// Moves to a city with taking into account if alien exausted its possible moves
				xwm, err = alien.sendToNextPosition(xwm)

				if err != nil {
					return nil, err
				}

				// Check if alien is trapped after the previous step
				if alien.isTrapped(xwm) {
					if !isInTrappedRecords(alien) {
						alien.accountTrapped()
					}
				}

				// check if alien exausted its moves
				if alien.moveCount == maxMovementsAllowed {
					alien.accountExausted()
				}
			}

			// Try to find opponent to fight with
			// Aliens that has status "trapped" and "exausted" also can fight
			xwm, _, err = alien.findOponent(xwm)
			if err != nil {
				return nil, err
			}
		}

		// Check if all aliens are destroyed before reaching maximum possible moves allowed
		if len(aliensReleased.destroyed) == len(aliensReleased.all) {
			fmt.Println("All aliens have been destroyed.")
			return xwm, nil
		}

		if len(aliensReleased.inOperation) == 0 {
			fmt.Println("No aliens able to operate left.")
			return xwm, nil
		}
	}

	return xwm, nil
}

// InvadeXWorld starts XWorld invasion
// Lands aliens across the cities and start their operational logic
func InvadeXWorld(filename string, aliensCount int) error {

	// Obtain data from input resource and return error
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	filename = filepath.Join(workingDir, "xworldmap", filename)
	xwmInitial, err := xWorldMap(filename)
	if err != nil {
		return err
	}

	// Create aliens and place them on their initial positions
	xwm := land(aliensCount, xwmInitial)

	// Start loop operation
	_, err = operate(xwm)

	fmt.Println("XWorld map status after invasion:")
	fmt.Printf("%d cities have been destroyed\n", len(xwm.Destroyed))
	fmt.Printf("%d cities still exist\n", len(xwm.Existing))

	fmt.Printf("Number of aliens able to operate: %d \n", len(aliensReleased.inOperation))
	fmt.Printf("Number of aliens trapped: %d \n", len(aliensReleased.trapped))
	fmt.Printf("Number of aliens destroyed: %d \n", len(aliensReleased.destroyed))
	fmt.Printf("Number of aliens exausted: %d \n", len(aliensReleased.exausted))

	err = cities.WriteXWorldCitiesData(xwm)

	return nil
}
