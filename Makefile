workshop-build:
	mdbook build

workshop-publish: workshop-build
	cd book && git init && git commit --allow-empty -m 'Update docs' && git checkout -b gh-pages && touch .nojekyll && git add . && git commit -am 'Update docs' && git push git@github.com:jbelmont/algorithms-workshop.git gh-pages --force

workshop-serve:
	mdbook serve 

workshop-book:
	open $(shell printf 'https://marcelbelmont.com/%s' $(shell git config remote.origin.url | sed 's/git.*\///g;s/\.git//g'))

workshop-github:
	open 'https://github.com/jbelmont/algorithms-workshop'

ifndef GO111MODULE
GO111MODULE = on
endif

ci-test:
	go test ./...