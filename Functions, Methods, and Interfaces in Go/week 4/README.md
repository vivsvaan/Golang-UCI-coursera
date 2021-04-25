# Peer-graded Assignment: Module 4 Activity

Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.


|  Animal  |  Food Eaten  |  Locomotion Method  |  Spoken Sound  |
|  ------  |  ----------  |  -----------------  |  ------------  |
|   Cow    |    Grass     |        Walk         |       Moo      |
|   Bird   |    Worms     |        Fly          |      Peep      |
|   Snake  |    Mice      |       Slither       |      Hsss      |

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”.  Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.
