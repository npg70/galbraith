
combined: build serve

hugo/content/people:
	mkdir -p ./hugo/content/tags
	mkdir -p ./hugo/content/people
	mkdir -p ./hugo/content/indexes
	mkdir -p ./hugo/content/lineages
	mkdir -p ./hugo/layouts/tags


build: hugo/content/people
	go run . -out hugo/content/people

test:
	go test .
serve:
	(cd hugo && hugo server -D)
clean:
	rm -rf hugo/public/*
	rm -rf hugo/content/people
	rm -rf hugo/content/indexes
	rm -rf hugo/content/lineages
	rm -rf hugo/content/tags

roots:
	go run . -root

