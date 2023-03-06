
combined: build serve

hugo/content/people:
	mkdir -p ./hugo/content/people
	mkdir -p ./hugo/content/indexes
	mkdir -p ./hugo/content/lineages

build: hugo/content/people
	go run . -out hugo/content/people

serve:
	(cd hugo && hugo server -D)
clean:
	rm -rf hugo/public/*
	rm -rf hugo/content/people
	rm -rf hugo/content/indexes
	rm -rf hugo/content/lineages

roots:
	go run . -root

