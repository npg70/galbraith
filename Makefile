
combined: build serve

hugo/content/people:
	mkdir -p ./hugo/content/people
	mkdir -p ./hugo/content/pages

build: hugo/content/people
	go run . -out hugo/content/people galbreath-james-1659-nielson

serve:
	(cd hugo && hugo server -D)
clean:
	rm -rf hugo/public/*
	rm -rf hugo/content/people/*

roots:
	go run . -root

