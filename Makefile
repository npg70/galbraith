
combined: build serve

hugo/content/people:
	mkdir -p ./hugo/content/people

build: hugo/content/people
	go run . -out hugo/content/people galbreath-james-1659-nielson
	go run . -out hugo/content/people galbreath-donald-1793-bell
	hugo build
serve:
	(cd hugo && hugo server -D)
clean:
	rm -rf hugo/public/*
	rm -rf hugo/content/people/*


