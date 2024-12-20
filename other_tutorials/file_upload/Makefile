install:
	rm -f ~/go/bin/file_upload
	go install -v
.PHONY: install

init:
	rm -f go.mod go.sum
	go mod init
	go mod tidy	
.PHONY: init

clean:
	rm -f ~/go/bin/file_upload
	rm -f go.mod
	rm -f go.sum

.PHONY: clean

