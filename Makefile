TAG=$(shell git rev-parse --short HEAD)

docker-build-bs:
	docker build --build-arg GITHUB_ACCESS_TOKEN="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" -t local/backserver -f backserver/Dockerfile  .

docker-build-web:
	cd web && npm run build
	cd web && docker build -t local/web -f Dockerfile .

docker-build: docker-build-bs docker-build-web

docker-push-bs-without-build:
	docker tag local/backserver 127.0.0.1:5000/backserver:$(TAG)
	docker push 127.0.0.1:5000/backserver:$(TAG)
	docker tag local/backserver 127.0.0.1:5000/backserver:latest
	docker push 127.0.0.1:5000/backserver:latest

docker-push-web-without-build:
	docker tag local/web 127.0.0.1:5000/web:$(TAG)
	docker push 127.0.0.1:5000/web:$(TAG)
	docker tag local/web 127.0.0.1:5000/web:latest
	docker push 127.0.0.1:5000/web:latest


docker-push-without-build: docker-push-bs-without-build docker-push-web-without-build

docker-push-bs: docker-build-bs docker-push-bs-without-build

docker-push-web: docker-build-web docker-push-web-without-build

docker-push: docker-push-bs docker-push-web

k8s-deploy: docker-push
	cat scripts/backserver-k8s.yml | sed 's/latest/$(TAG)/' > /tmp/$(TAG).yml
	kubectl --kubeconfig ~/.kube/config apply -f /tmp/$(TAG).yml --all
	cat scripts/web-k8s.yml | sed 's/latest/$(TAG)/' > /tmp/$(TAG).yml
	kubectl --kubeconfig ~/.kube/config apply -f /tmp/$(TAG).yml --all

k8s-deploy-without-docker-push:
	cat scripts/backserver-k8s.yml | sed 's/latest/$(TAG)/' > /tmp/$(TAG).yml
	kubectl --kubeconfig ~/.kube/config apply -f /tmp/$(TAG).yml --all
	cat scripts/web-k8s.yml | sed 's/latest/$(TAG)/' > /tmp/$(TAG).yml
	kubectl --kubeconfig ~/.kube/config apply -f /tmp/$(TAG).yml --all