
combined: build serve

hugo/content/people:
	mkdir -p ./hugo/content/people
	mkdir -p ./hugo/content/indexes
	mkdir -p ./hugo/content/lineages
	mkdir -p ./hugo/content/sources


build: hugo/content/people
	go run . -out hugo/static

test:
	go test .
serve:
	(cd hugo && hugo server -D)
clean:
	rm -f .git/index.lock
	rm -rf hugo/public/*
	rm -rf hugo/static/people hugo/content/people
	rm -rf hugo/static/indexes hugo/content/indexes
	rm -rf hugo/static/tags hugo/content/tags
	rm -rf hugo/static/sources hugo/content/sources

roots:
	go run . -root

lint:
	golangci-lint run .
env:
	brew upgrade
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
