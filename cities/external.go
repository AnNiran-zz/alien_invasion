package cities

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadXWorldCitiesData reads input file data line by line and saves data inside cities.XWorldMap struct
// which will be used as a directions map for aliens movements
func (xwm *XWorldMap) ReadXWorldCitiesData(filename string) error {
	fmt.Println("Reading XWorld cities data and obtaining directions map ...")

	posDirections := make(map[string]string)

	// Add folder name to path and read data from file
	inputData, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer inputData.Close()

	scanner := bufio.NewScanner(inputData)

	// Read data dfrom file
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		cityName := data[0]
		routes := data[1:]

		for _, route := range routes {
			rsplit := strings.Split(route, "=")

			direction := rsplit[0]
			toCity := rsplit[1]

			if direction != "north" && direction != "south" && direction != "west" && direction != "east" {
				return fmt.Errorf("input data format does not match expected")
			}

			posDirections[direction] = toCity
		}

		// Add read line data as a new city to the map
		err = xwm.AddCity(cityName, posDirections)
		if err != nil {
			return err
		}

		// Reset used map
		posDirections = make(map[string]string)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// WriteXWorldCitiesData saves updated XWorld map after the alien invasion finished
func WriteXWorldCitiesData(xwm *XWorldMap) error {
	fmt.Println("Writing XWorld map conditions data in xworld-map-output.txt file")

	outputfile, err := os.OpenFile("xworldmap/xworld-map-output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Reset file content after the previous run
	outputfile.Truncate(0)
	if len(xwm.Existing) < 1 {
		fmt.Println("No cities left from XWorld. No data will be output.")
		return nil
	}

	var outputData []string
	var line string
	for city, directions := range xwm.Existing {
		line = fmt.Sprintf("%s", city)

		for direction, destination := range directions {
			line += fmt.Sprintf(" %s=%s", direction, destination)
		}
		outputData = append(outputData, line)
		line = ""
	}

	writer := bufio.NewWriter(outputfile)

	for _, data := range outputData {
		_, _ = writer.WriteString(data + "\n")
	}

	writer.Flush()
	outputfile.Close()

	fmt.Println("XWorld updated map exported")

	return nil
}
