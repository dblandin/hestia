.PHONY: build dep test

IMAGE_NAME := codeclimate/hestia

image:
	docker build -t $(IMAGE_NAME) .

build:
	docker run -it --rm  \
	  --volume "$(CURDIR):/go/src/github.com/codeclimate/hestia" \
	  $(IMAGE_NAME) scripts/build

package:
	zip -j dist/package.zip dist/*

test:
	docker run -it --rm  \
	  --volume "$(CURDIR):/go/src/github.com/codeclimate/hestia" \
	  $(IMAGE_NAME) scripts/test

release:
	aws s3 cp \
	  --acl private \
	  dist/package.zip \
	  "s3://${TF_VAR_release_s3_bucket}/${TF_VAR_release_s3_key}"
