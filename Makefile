# Makefile for Serverless Go

.PHONY: setup \
	precondition-aws \


setup:
	go mod download
	npm install

precondition-aws:
	aws sts get-caller-identity

package:
	go version
	sh scripts/package.sh

	serverless package

deploy: package precondition-aws
	serverless deploy --force

remove: precondition-aws
	serverless remove