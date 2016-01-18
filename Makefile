.PHONY: clean

all:
	go build -o build/tailor ./*.go

install: 
	go install -o tailor ./*.go

clean:
	rm build/tailor
