run:
	go run main.go

wind:
	./tailwindcss -i ./src/input.css -o ./static/css/tw.css --watch

build:
	go build main.go