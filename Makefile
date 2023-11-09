compile-proto:
	protoc -Iproto --go_out=. --go_opt=module=github.com/Sokol111/category-service --go-grpc_out=. --go-grpc_opt=module=github.com/Sokol111/category-service proto/*.proto

generate-mocks: compile-proto
	mockery

test: generate-mocks
	go test ./... -v -cover

build-server: test
	go build -o category-service ./cmd

build-docker-image: test
	docker build -t sokol111/category-service:latest .

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

run-server: build-server
	./category-service

push-image-to-dockerhub:
	docker push sokol111/category-service:latest

kubectl-apply:
	kubectl apply -f ./.k8s/config.yml
	kubectl get all -o wide
	minikube service category-service-node-port-service --url

kubectl-get-all:
	kubectl get all -o wide

kubectl-delete:
	kubectl delete -f ./.k8s/config.yml
	kubectl get deployments

kubectl-get-pods:
	kubectl get pods

kubectl-show-logs:
	kubectl logs -f -l app=category-service-label --all-containers --prefix=true --timestamps=true