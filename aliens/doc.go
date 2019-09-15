/*
Package aliens manage the creation and all changes in status and position for all aliens

Aliens number is landed on the current XWorld map with status "inoperation"
After that each alien starts to move randomly across the cities, using possible routes out of the city
it is located at each step

When other aliens are encountered in a city, it initiates a fight with one of them
(chosen randomly if potential opponents are more than 1)
and that action destorys the alien, all the other aliens, its opponent and the city
After that the map status is updated - by removing all routes to the destroyed city from other cities maps

If there are no other aliens encountered in a city - the alien moves on to the next city

At each iteration - XWorld map is updated, and checked if an alien is trapped after the last loop

operate() is the function that manages this process
Maximum allowed iterations are 10 000
If all aliens are trapped or destroyed before iterations reach that number - the loop stops
After the loop stops - the program outputs the current situation after the invasion:
how many cities are destroyed and how many still exist
how many aliens are destoyed and exausted or trapped

The updated XWorld map is output in a file: xworld-map-oputput.txt and the program end
and it holds the cities resulting data
*/
package aliens
