.PHONY: clean assets

all: assets darwin

assets:
	cd assets_src && browserify tailor.js -t babelify -o ../assets/tailor.js && uglifyjs -c -o ../assets/tailor.min.js ../assets/tailor.js
	cleancss assets_src/tailor.css -o assets/tailor.css
	cp assets_src/*.html assets/
	go-bindata -ignore node_modules -ignore package.json -pkg main -o static.go assets/

linux:
	GOOS=linux GOARCH=amd64 go build -o build/tailor ./*.go

windows:
	GOOS=windows GOARCH=amd64 go build -o build/tailor.exe ./*.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/tailor ./*.go

install: 
	go install -o tailor ./*.go

clean:
	rm build/tailor
