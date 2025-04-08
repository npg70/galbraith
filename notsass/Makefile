
build:
	go run . -input styles.css > junk.txt
	cp junk.txt ../hugo/static/notbootstrap.css

lint:
	go run . -input styles.css > junk.txt
	npx stylelint --config ./.stylelintrc.json junk.txt
