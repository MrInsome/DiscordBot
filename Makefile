build:
	go build -o ./build/discordBot ./cmd/main.go

run: build
	./build/discordBot