
combined: serve

build:
	go run . -out hugo/static 

test:
	go test .

serve:
	go run . -out hugo/static -serve

clean:
	rm -f .git/index.lock
	rm -rf hugo/public/*
	rm -rf hugo/static/people
	rm -rf hugo/static/indexes
	rm -rf hugo/static/tags
	rm -rf hugo/static/sources

roots:
	go run . -root

lint:
	golangci-lint run .

env:
	brew upgrade
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
