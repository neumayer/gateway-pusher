dist:
	mkdir dist
	GOOS=darwin GOARCH=amd64 go build -o dist/gateway-pusher-darwin-amd64 
	go build -o dist/gateway-pusher-linux-amd64 

clean:
	rm -rf dist

