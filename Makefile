default:
	go build hgen.go

install: default
	cp hgen $(HOME)/bin/

clean:
	rm -rf hgen
