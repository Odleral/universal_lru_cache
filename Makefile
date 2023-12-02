run:
	go run main.go

test:
	go test ./app

bench:
	go test -bench=. -benchtime=100x -benchmem ./app
