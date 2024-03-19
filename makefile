run:
	gofmt -w .
	templ fmt .
	templ generate
	go run .
