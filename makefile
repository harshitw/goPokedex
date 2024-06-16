build :
	@go build -o bin/pokedexCli
run : build
	@./bin/pokedexCli
test :
	@go test -v ./...