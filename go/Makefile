build:
	# outFile: bin/main, buildTarget: cmd/main.go
	go build -o bin/main cmd/main.go

test:
	ENV=test go test --cover ./...

generate:
	rm -rf mock/
	go generate ./...