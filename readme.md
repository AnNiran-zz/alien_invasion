Alien Invasion
==============

The project represents a simulation of an alien invasion on a wolrd called X

The main.go file is used to start the invasion
The program uses a filename inside xworldmap/ folder to read data from

the data should be in a specific format - example in cities/doc.go 

Program expects command line inputs after starting - filename to read from and number of aliens
to be created for the invasion

Invasion starts with creation of the specified number of aliens and placing them at random
places across the map of XWorld

Each alien has its own status depending on its condition, that may be:
* inoperation - means the alien can move further and is not trapped
* exausted - alien has reached the maximum alloed number of moves and cannot move further
* trapped - alien cannot move further because there are no existing routes of the citiy, it is located in 
* destroyed - alien is destroyed and no longer participate in any process

Initial status of each alien is "inoperation"

Changes to statuses "trapped" and "exausted" can be applied only on alien that has status "inoperation"
The purpose is to account the initial condition that followed to the unablity of the alien to move

Changes to status "destroyed" can be applied to all aliens, regardless of their current status

Cache data for all aliens and their statuses is kept and updated with program running

XWorld map contains records about cities that currently have aliens placed in them, still exist or are destroyed 
and is updated accordingly as well

The program stops when all aliens able to move have made the maximum allowed steps
or all have been destroyed

After invasion finishes - the program prints a report for the cities and aliens current conditions 
and outputs the existing cities data in a file, using the same format as the input one

