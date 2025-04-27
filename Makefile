# Variables
DOCKER_USERNAME=vijayk93
SERVICE_GENERATE_IMAGE=$(DOCKER_USERNAME)/shortitsh:service-generate
SERVICE_REDIRECT_IMAGE=$(DOCKER_USERNAME)/shortitsh:service-redirect
HELM_RELEASE_NAME=shortitsh
NAMESPACE=shortit

# Build service-generate
build-service-generate:
	docker build -t $(SERVICE_GENERATE_IMAGE) ./service-generate

# Build service-redirect
build-service-redirect:
	docker build -t $(SERVICE_REDIRECT_IMAGE) ./service-redirect

# Push service-generate image
push-service-generate:
	docker push $(SERVICE_GENERATE_IMAGE)

# Push service-redirect image
push-service-redirect:
	docker push $(SERVICE_REDIRECT_IMAGE)

# Deploy using Helm
deploy:
	helm upgrade --install $(HELM_RELEASE_NAME) ./deployments -f ./deployments/values.yaml --namespace $(NAMESPACE)

# Clean up the Helm release
clean:
	helm uninstall $(HELM_RELEASE_NAME) --namespace $(NAMESPACE)

# Full pipeline: Build, Push, and Deploy
all: build-service-generate build-service-redirect push-service-generate push-service-redirect deploy
