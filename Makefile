
combined: serve

build: 
	cd notsass && make
	mkdir -p public
	cp -rf static/* public/
	go run .

test:
	go test .


serve:
	cd notsass && make
	mkdir -p public
	cp -rf static/* public/
	go run . -serve

clean:
	go clean -i -r -cache -testcache
	rm -f .git/index.lock
	rm -rf public
	rm -f static/fulltext.json

lint:
	go mod tidy
	gofmt -w -s *.go
	golangci-lint run .

env:
	brew upgrade
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
