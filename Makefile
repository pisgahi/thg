wind:
	./tailwindcss -i ./src/input.css -o ./static/css/tw.css --watch

air: 
	air

run:
	go run main.go

build:
	go build main.go
