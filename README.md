# MySafeApp

## Overview

MySafeApp is a project designed for educational purposes, specifically to teach DevOps principles using k3s and ArgoCD. This project demonstrates the integration of these technologies to create, manage, and deploy a containerized application in a lightweight Kubernetes environment.

## Technology Stack

- **Frontend:** Vue.js
- **Backend:** Go (Golang)
- **Database:** MySQL
- **Containerization:** Docker
- **Orchestration:** k3s (Lightweight Kubernetes)
- **CI/CD:** ArgoCD

## Project Structure

The project is structured into three main components:

1. **Frontend:** The user interface built with Vue.js.
2. **Backend:** The server-side logic built with Go.
3. **Database:** The MySQL database to store application data.

## Getting Started

### Prerequisites

- Docker installed
- k3s installed and running
- ArgoCD installed and configured
- Access to Docker Hub or a container registry

### Building the Docker Images

Build and push the Docker images for the frontend and backend components:

```sh
# Build the frontend image
cd frontend
docker build -t <your-dockerhub-username>/mysafeapp-frontend:v1.0.0 .
docker push <your-dockerhub-username>/mysafeapp-frontend:v1.0.0

# Build the backend image
cd ../backend
docker build -t <your-dockerhub-username>/mysafeapp-backend:v1.0.0 .
docker push <your-dockerhub-username>/mysafeapp-backend:v1.0.0


### Deploying to k3s
Apply the Kubernetes manifests to deploy the application to your k3s cluster:

sh
kubectl apply -f k8s/mysql-deployment.yaml
kubectl apply -f k8s/backend-deployment.yaml
kubectl apply -f k8s/frontend-deployment.yaml
### Configuring ArgoCD
Use ArgoCD to manage and deploy the application:

Install ArgoCD and access the ArgoCD UI.
Create an ArgoCD application pointing to your GitHub repository.
Sync the application to deploy it to your k3s cluster.
Educational Goals
This project aims to teach the following DevOps practices:

Containerization: Building and managing Docker containers for frontend and backend services.
Kubernetes: Deploying and managing applications in a k3s cluster, a lightweight Kubernetes distribution.
CI/CD: Setting up continuous integration and continuous deployment pipelines using ArgoCD.
Automation: Automating deployment and scaling processes using Kubernetes and ArgoCD.
Conclusion
MySafeApp is a practical example of a modern application deployment using lightweight Kubernetes (k3s) and a powerful GitOps tool (ArgoCD). It is an excellent resource for anyone looking to gain hands-on experience with these technologies and improve their DevOps skills.

### Contact
For any questions or suggestions, feel free to open an issue or contact the maintainer.
