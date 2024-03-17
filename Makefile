
combined: serve

hugo/public:
	mkdir -p hugo/public

build: hugo/public
	cp -rf hugo/static/* hugo/public/
	go run . -out hugo/public

test:
	go test .

serve: hugo/public
	cp -rf hugo/static/* hugo/public/
	go run . -out hugo/public -serve

clean:
	rm -f .git/index.lock
	rm -rf hugo/public

lint:
	golangci-lint run .

env:
	brew upgrade
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
