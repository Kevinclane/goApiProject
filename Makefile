build: 
	GOOS=darwin GOARCH=amd64 go build -o api main.go

clean:
	rm api