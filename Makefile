# Variables
APP_NAME = url-shortener
GO_VERSION = 1.19
CLUSTER_NAME = url-shortener-cluster
CLUSTER_CONFIG = cluster/kind-cluster-config.yaml
DOCKER_IMAGE = cyrilbaah/$(APP_NAME):latest
K8S_MANIFEST = k8s/deployment.yaml

# Help documentation
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help            Show this help message"
	@echo "  build           Build the Go application"
	@echo "  run             Run the application locally"
	@echo "  docker-build    Build the Docker image"
	@echo "  docker-run      Run the application in a Docker container"
	@echo "  kind-create     Create a Kind cluster"
	@echo "  kind-delete     Delete the Kind cluster"
	@echo "  k8s-deploy      Deploy the application to Kubernetes"
	@echo "  port-forward    Forward service port to localhost"
	@echo "  clean           Remove all stopped containers and images"

# Build the Go application
build:
	@echo "Building $(APP_NAME)..."
	@go mod tidy
	@go build -o $(APP_NAME)
	@echo "Build complete!"

# Run the application locally
run: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME)

# Build the Docker image
docker-build:
	@echo "Building Docker image $(DOCKER_IMAGE)..."
	@docker build -t $(DOCKER_IMAGE) .
	@echo "Docker build complete!"

# Run the Docker container
docker-run:
	@echo "Running Docker container for $(DOCKER_IMAGE)..."
	@docker run -p 8080:8080 $(DOCKER_IMAGE)

# Create a Kind cluster
kind-create:
	@echo "Creating Kind cluster $(CLUSTER_NAME)..."
	@kind create cluster --name $(CLUSTER_NAME) --config $(CLUSTER_CONFIG)
	@echo "Kind cluster created!"

# Deploy to Kubernetes
k8s-deploy:
	@echo "Deploying to Kubernetes..."
	@kubectl apply -f $(K8S_MANIFEST)
	@echo "Deployment complete!"

# Delete the Kind cluster
kind-delete:
	@echo "Deleting Kind cluster $(CLUSTER_NAME)..."
	@kind delete cluster --name $(CLUSTER_NAME)
	@echo "Kind cluster deleted!"

# Make service accessible
port-forward-app:
	@echo "Forwarding service port to localhost..."
	@kubectl port-forward service/url-shortener 8080:80
	@echo "Port forwarding active at http://localhost:8080"

# Make prometheus accessible
port-forward-prometheus:
	@echo "Forwarding prometheus port to localhost..."
	@kubectl port-forward service/prometheus 9090:9090
	@echo "Port forwarding active at http://localhost:9090"

# Make grafana accessible
port-forward-grafana:
	@echo "Forwarding grafana port to localhost..."
	@kubectl port-forward service/grafana 3000:3000
	@echo "Port forwarding active at http://localhost:3000"


# Clean up all stopped containers and images
clean:
	@echo "Cleaning up Docker containers and images..."
	@docker stop $(shell docker ps -aq) || true
	@docker rm $(shell docker ps -aq) || true
	@docker rmi $(shell docker images -aq) || true
	@echo "Cleanup complete!"

.PHONY: help build run docker-build docker-run kind-create kind-delete k8s-deploy port-forward clean