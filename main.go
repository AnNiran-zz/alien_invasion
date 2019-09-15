package main

import (
	"alien_invasion/aliens"
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// main starts the program and ask for user input for
// file to read from and number of aliens to be landed in XWorld map
func main() {
	fmt.Println("Welcome to XWorld Alien invasion!")

	file, err := readFilename()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Please try again")

		os.Exit(1)
	}

	aliensCount, err := readAliensCount()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Please try again")

		os.Exit(1)
	}

	// Start invasion
	fmt.Printf("Starting alien invasion of XWorld with %d aliens\n", aliensCount)
	err = aliens.InvadeXWorld(file, aliensCount)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Good bye!")
	os.Exit(1)
}

// Read input filename to obtain map cities data from
func readFilename() (string, error) {
	fmt.Println("Please provide input filename to read XWorld cities data from.")
	fmt.Println("File must be located inside xworldmap/ project folder:")

	value, err := reader.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("Error reading input")
	}

	file := strings.TrimSuffix(value, "\n")

	location := filepath.Join("xworldmap/", file)
	// check if provided filename exists in xworldmap folder
	if _, err := os.Stat(location); os.IsNotExist(err) {
		fmt.Println("Provided file does not exists in xworldmap/ folder")
		fmt.Println("Please provide a valid file")
		fmt.Println("*********")

		return "", errors.New("File does not exist")
	}

	fmt.Printf("Filename entered to read from is %s\n", file)

	return file, nil
}

// Read input number of aliens to be released during the invasion
func readAliensCount() (int, error) {
	fmt.Println("Please provide number of aliens to be unleashed:")

	value, err := reader.ReadString('\n')

	if err != nil {
		return 0, fmt.Errorf("Error reading input")
	}

	value = strings.TrimSuffix(value, "\n")

	value = strings.TrimSuffix(value, "\n")

	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("Unable to convert provided value: %s to integer", value)
	}

	fmt.Printf("Number of aliens to be released is: %d \n", number)

	return number, nil
}
