# TicTacToe
A two player game

#### Install
git clone https://github.com/vaibhawj/tictactoe.git $GOPATH/src/github.com/vaibhawj/tictactoe

#### Unit tests
cd into project root
go test ./... -cover -coverprofile coverage.out; go tool cover -func coverage.out

#### Run
go run github.com/vaibhawj/tictactoe