alias r := run
alias b := build
alias m := model
alias gen := generate

# install the dependencies
init:
	go get
	go install github.com/cosmtrek/air@1.29.0
# run the server with hot reload
run:
	air
# build the server
build:
	go build -o ./tmp/main ./cmd/server
# create a ent schema
model NAME:
	go run -mod=mod entgo.io/ent/cmd/ent new {{ NAME }}
# generate the ent schema
generate:
	go generate ./ent
