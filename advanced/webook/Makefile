.PHONY: docker
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook .
	@docker rmi -f gevin/webook:0.0.1 || true
	@docker build -t gevin/webook:0.0.1 .