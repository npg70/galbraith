
combined: serve

build: 
	cd notsass && make
	mkdir -p public
	cp -rf static/* public/
	go run .

test:
	go test .

# NOTE: for github, the deployed files are at /galbraith/...
#   not /
# 
#
serve:
	cd notsass && make
	mkdir -p public/galbraith
	cp -rf static/* public/galbraith/
	go run . -out public/galbraith
	python3 -m http.server --directory public --bind 127.0.0.1 1313  

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
