## Tic Tac Toe - Project-0

Tictactoe is a CLI application that takes in user input and plays tictactoe. You can play by yourself against Ai, You can play against another player or you cannot play at all and watch the outcome of Ai vs Ai. 

## Requirements

Go is a a requirement and is available for download at https://golang.org

Windows users will need to use bash to build this project. Bash is incorporated into git and is available at https://gitforwindows.org/ 

## Installation

In an open terminal / bash, download the project with command: `go get -u github.com/JesseyBland/project-0`

Now still in terminal / bash, navigate to the `project-0` installed directory $user/go/src/JesseyBland/project-0 then 
run both of these commands:

CLI
```bash
go build ./cmd/tictactoe
```
HTTP server
```bash
go build ./cmd/tictactoed
```

## Command Line Interface

To start the Tic Tac Toe Cli application navigate in terminal/bash to the tictactoe program built in the $user/go/src/JesseyBland/project-0 directory and run it:
```bash
./tictactoe
```

Command-line flag arguments can be called to change the game mode.

Command-line flags:

**-tp**
Two Player mode

**-av**
Ai vs Ai mode

**-g1,g2,g3,g4,g5,g6**
Board design options 7 counting the default.

Example:
```bash
$user/go/src/JesseyBland/project-0
./tictactoe -tp
```
## Http server 

The server runs off Localhost:8080 includes directories /tt1 /ttt2 /ttt3. These have html file dependencies included inside the package tictactcoed in the folder called web. The http server must be run from the project directory for the html templates to function. 

To start the server navigate in terminal/bash to the project file where tictactoed was built and run the program.
Now look for the Server status.

Example:
```bash
$user/go/src/JesseyBland/project-0
$./tictactoed 
Server Status:Listening Host:localhost Port:8080 
```
Once the server has started, navigate in your browser to URL localhost:8080

3 available play modes 

**Player vs AI**

**Player vs Player**

**Ai vs Ai**

## Features Road map

- [x] User can play against AI
- [x] 2 player flag option
- [x] Respond to username
- [x] Add http browser play
- [x] Cosmetic options
- [ ] Add Reverse Proxy to http server
