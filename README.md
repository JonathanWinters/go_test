# GO and Postgres Test


## Use Instructions:
1. Create a docker instance of https://hub.docker.com/_/postgres and name it "pg-test" with a password of "secret"
2. cd into `cmd/server`, type `go run main.go` and press enter
3. cd into `scripts/endpoints` and type `./submit.sh`, give it execute commands if needed using `chmod +x`, then press enter
4. type `./move.sh {your move [0, 1, 2, 3, 4]}` in order to move your player. Look at the move.json file inside the tmp folder that gets created to see updates.

## Test Overview:
For this test you will be writing a small web server in Golang and using Postgres as a database.

For each part, please be sure to explain your solution, the choices you’ve made, and note how long it took you to complete. Also, send me your final code and database schema in an email with your write-up along with all instructions necessary to build and run your program.

Do not be afraid to reach out if any portion of the test is confusing or unclear!

### Part 1: Creating Levels
You are building a game where users can submit levels they have created to be played by other players. Create a /submit endpoint that takes in a level in JSON format. Submitted levels should be stored in a Postgres database, and if successful the endpoint should return an id by which the submitted levels can be referenced.

Levels are arrays of arrays of numbers, where the position in the arrays represents the (X, Y) position in the level and the number represents the object at that location. 

Use the following mapping:
```
0 - open tile
1 - wall
2 - pit trap. Can be moved through but player would take 1 damage
3 - arrow trap. Can be moved through but player would take 2 damage
4 - player starting position (will be open after they leave it)
```

An example Level JSON might look like this:
```
[
[1,1,1,1,0,1,1,1],
[1,0,0,0,0,0,0,1],
[1,0,1,1,1,3,1,1],
[1,0,0,0,1,0,2,1],
[1,1,1,0,1,1,0,1],
[1,0,0,0,1,0,0,1],
[1,0,1,1,1,0,1,1],
[1,0,0,4,0,0,0,1],
[1,1,1,1,1,1,1,1]
]
```

### Part 2: Validation & Testing
We want a little more validation of the maps that the players are storing. Upon submission the user should receive a descriptive error message if their map fails any of the following validation checks:
● Maps must be rectangular
● Maps may not be larger than 100 in any dimension.
● Map spaces may not use values other than the numbers 0-4 above.

In addition to the above, are there validation steps you would suggest? You don’t need to program them, but please discuss any additional validation you would do.

Also, add a few examples of what you consider good tests. Don’t worry about being extensive or about test coverage. Just supply a sample of ideas you think might be useful.

### Part 3: Moving the Player
Add an endpoint for moving the player around a given map. Don’t create a new copy of the map per user, multiple users should be able to control the same player’s moves around the same map. When two users are moving the player at the same time, they should not overwrite each other's moves, but instead should have their moves applied one after the other.
This endpoint will take the id of the map and a number representing the direction to move the player using the following format:
```
0 - left
1 - up
2 - right
3 - down
```
This endpoint will return a JSON that is the new state of the map in the format from part 1. 

Assume the player has 4 hit points. If a player would die from taking too much damage from moving onto traps, restart the player where his start position was in the initial submission of the map.

### Part 4 (Optional): Minimum Survivable Path
Now assume the player has 4 hit points. We want to calculate the minimum survivable path. The minimum survivable path is the path that gets to an exit of the maze without the player dying in the minimum number of moves. For instance, the example map has two primary paths to the exit: one which takes no damage and the distance is 16, and one where the player takes 3 damage and the distance is 12. In this case the minimum survivable path to the exit is the path of length 12. If the pit trap was replaced by an arrow trap, the player would die if they tried to take the 12-move path, and thus the 16 move path would be the minimum survivable path. Describe an algorithm for finding the minimum survivable path. Describe its run time using Big O notation. You do not need to implement the algorithm.