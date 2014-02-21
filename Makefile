default:
	go build hgen.go

install: default
	cp hgen $(HOME)/bin/

fmt:
	go fmt

clean:
	rm -rf hgen
