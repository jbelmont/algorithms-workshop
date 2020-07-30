workshop-build:
	mdbook build

workshop-publish: workshop-build
	cd book && git init && git commit --allow-empty -m 'Update docs' && git checkout -b gh-pages && touch .nojekyll && git add . && git commit -am 'Update docs' && git push git@github.com:jbelmont/algorithms-workshop.git gh-pages --force