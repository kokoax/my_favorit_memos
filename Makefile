build:
	go build -o my_favorit_memos main.go
test:
	go test -v ./...
run: build
	./my_favorit_memos
deploy: build
	cp ./my_favorit_memos ~/.bin/my_favorit_memos
