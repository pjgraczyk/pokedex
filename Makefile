.PHONY: all build run test clean checkapi fmt format

all: build

build:
	go build -o pokedexcli .

run:
	go run .

test:
	go test -v ./...

clean:
	rm -f pokedexcli

fmt:
	go fmt ./...

format: fmt

checkapi:
	http GET https://pokeapi.co/api/v2/pokemon/$(or $(POKEMON),ditto) | jq | bat --pager=less --force-colorization -l json
