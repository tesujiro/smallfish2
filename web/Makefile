MODULE=smallfish2web
GOOS=linux
REPO="localhost:5000"
POD=sf2-web-rs-

build:
	GOOS=$(GOOS) go build -o $(MODULE).$(GOOS)
test:
	go test -v .

.PHONY: container
container: container_build container_deploy

.PHONY: container_build
container_build:
	docker build -t $(MODULE) .

.PHONY: container_deploy
container_deploy:
	docker tag $(MODULE) $(REPO)/$(MODULE)
	docker push $(REPO)/$(MODULE)
	docker rmi $(REPO)/$(MODULE)
	docker pull $(REPO)/$(MODULE)

.PHONY: pod_reset
pod_reset:
	for p in `kubectl get pod | grep $(POD) | awk '{print $$1}'`;do\
		kubectl delete pod $$p;\
	done
	docker rmi $$(docker images -a --filter "dangling=true" -q) -f

.PHONY: apply
apply:
	kubectl apply -f deploy.yaml

SSL_KEY=ssl/server.key
SSL_CSR=ssl/server.csr
SSL_CRT=ssl/server.crt
.PHONY: ssl_certificate
ssl_certificate:
	-mkdir ssl
	openssl genrsa 2048 > $(SSL_KEY)
	openssl req -new -key $(SSL_KEY) > $(SSL_CSR)
	openssl x509 -days 3650 -req -signkey $(SSL_KEY) < $(SSL_CSR) > $(SSL_CRT)
