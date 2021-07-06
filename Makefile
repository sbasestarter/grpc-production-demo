TAG=$(shell git rev-parse --short HEAD)

daemon:
	# start our program as a background process (daemon).
	nohup go run app/demo.go &

# 启动 server
run:
	go run app/demo.go

testing:
	go test -v -count 1 ./...
	golangci-lint run -v
	kube-linter lint k8s-deployment.yml

build:
	CGO_ENABLED=0 GOOS=linux go build -o Niffler -v -ldflags "-w" app/demo.go

docker-build-bs:
	docker build --build-arg GITHUB_ACCESS_TOKEN="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" -t local/backserver -f backserver/Dockerfile  .

docker-push-without-build:
	docker tag local/backserver 127.0.0.1:5000/backserver:$(TAG)
	docker push 127.0.0.1:5000/backserver:$(TAG)
	docker tag local/backserver 127.0.0.1:5000/backserver:latest
	docker push 127.0.0.1:5000/backserver:latest

docker-push: docker-build-bs docker-push-without-build

k8s-deploy: docker-push
	cat scripts/backserver-k8s.yml | sed 's/latest/$(TAG)/' > /tmp/$(TAG).yml
	kubectl --kubeconfig ~/.kube/config apply -f /tmp/$(TAG).yml --all

k8s-deploy-without-docker-push:
	cat scripts/backserver-k8s.yml | sed 's/latest/$(TAG)/' > /tmp/$(TAG).yml
	kubectl --kubeconfig ~/.kube/config apply -f /tmp/$(TAG).yml --all