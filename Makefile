build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/demo-k8s

docker:
	docker build -t github.com/as-polyakov/demo-k8s .

test:
	go test -v .

