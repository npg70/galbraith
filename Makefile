
combined: serve

hugo/public:
	mkdir -p hugo/public

build: hugo/public hugo/static/darker.css
	cp -rf hugo/static/* hugo/public/
	go run . -out hugo/public

test:
	go test .

hugo/static/darker.css: notsass/tailwind-dark.css
	cd notsass && go run . -input 'tailwind-dark.css' > ../hugo/static/darker.css

serve: hugo/public hugo/static/darker.css
	cp -rf hugo/static/* hugo/public/
	go run . -out hugo/public -serve

clean:
	go clean -i -r -cache -testcache
	rm -f .git/index.lock
	rm -rf hugo/public
	rm -f hugo/static/fulltext.json

lint:
	golangci-lint run .

env:
	brew upgrade
	brew install golangci-lint
	go install golang.org/x/tools/cmd/goimports@latest
