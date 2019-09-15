/*
Package cities implements functionality for manipulating data
about currently existing cities

Data is obtained from input source via command line argument and saved in a structure
that represent the XWorld map of cities

Data file is expected to contain direction details about currently existing cities
in a specific format:

Example:
CityA south=CityE east=CityB
CityB south=CityF west=CityA east=CityC
CityC south=CityG west=CityB east=CityD
CityD south=CityH west=CityC

The structure is continuously updated while the program processes, using a pointer to it
Updates include cities being destroyed, removing of all routes to a destroyed city,
account of an alien arrival or leave in the city

The package functions are used by package aliens
*/
package cities
