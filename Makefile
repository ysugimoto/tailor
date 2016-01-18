.PHONY: clean

all: darwin

linux:
	GOOS=linux GOARCH=amd64 go build -o build/tailor ./*.go

windows:
	GOOS=windows GOARCH=amd64 go build -o build/tailor ./*.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/tailor ./*.go

install: 
	go install -o tailor ./*.go

clean:
	rm build/tailor
