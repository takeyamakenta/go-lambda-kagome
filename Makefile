PHONY: build pack
build:
	rm -rf dist && mkdir dist && GOOS=linux GOARCH=amd64 go build -o ./dist/main main.go

pack:
	zip -r ./main.zip ./dist/