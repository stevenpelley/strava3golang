all: | openapi-generator readme

openapi-generator:
	docker run --rm \
	-v $(PWD):/local/go openapitools/openapi-generator-cli generate \
	--output /local/go \
	--input-spec https://developers.strava.com/swagger/swagger.json \
	--generator-name go \
	-p enumClassPrefix=true \
	--skip-validate-spec \
	--package-name strava3golang \
	--git-host github.com \
	--git-user-id stevenpelley \
	--git-repo-id strava3golang

readme:
	cat README_HEADER.md README.md > README.md.tmp; mv README.md.tmp README.md

gitignore:
	cat .gitignore_header .gitignore > .gitignore.tmp; mv .gitignore.tmp .gitignore

clean:
	find . ! -name 'Makefile' ! -name 'README.md' ! -name 'LICENSE' ! -name '.git' ! -name '.' ! -name 'README_HEADER.md' -maxdepth 1 -exec rm -rf {} +
