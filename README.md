# Project Overview

This project demonstrates the basic workings of Docker and Kubernetes, focusing on:  
1. A simple Hello-world Docker container.  
2. A Python Flask ping-pong service that responds to "ping".

---

## Setting Up Docker and Kubernetes

Docker was pre-installed on my machine. I enabled Kubernetes in Docker to deploy and manage containerized applications.

---

## Hello World

### Running Locally with Docker
I pulled the `hello-world` Docker image from Docker Hub and ran it in a Docker container on my local machine. This service prints "Hello, Docker!" to the Terminal.

### Deploying with Kubernetes
I deployed the `hello-world` Docker container using Kubernetes. The `hello-world-docker` folder contains a `hello-world.yaml` file, which defines the application and its deployment to a Kubernetes cluster.

---

## Breakdown of `hello-world.yaml`

- **API Version**:  
  The `apiVersion` is `v1`, which indicates the use of a stable Kubernetes API.

- **Kind**:  
  The `kind` is `Pod`, which specifies that the deployment will run the `hello-world` container inside a Kubernetes Pod.

- **Metadata**:  
  - `name`: The name of the Pod, which is set to `hello-world`.

- **Spec**:  
  - `containers`: Contains the configuration for the container:
    - `name`: The name of the container, which is `hello-world`.
    - `image`: Specifies the Docker image to use, which is `hello-world`.

---

## Steps to Run on Terminal

### Running Locally with Docker
1. **Pull the Docker Image**:  
   Downloaded the `hello-world` image from Docker Hub using the following command:  
   ```bash
   docker pull hello-world
   ```
2. **Run the Docker Container**:
   Started the container locally to see the `"Hello from Docker!"` message:
  ```bash
  docker run hello-world
```
### Deploying with Kubernetes
1. **Apply the Kubernetes Configuration**:
  The hello-world.yaml file contains deployment specifications. Used it to deploy the application to a Kubernetes cluster:
```bash
kubectl apply -f hello-world-docker/hello-world.yaml
```

-----------------
## Ping Pong Service in Python

## Overview
This project demonstrates a Python Flask application that listens on port `5000` and responds to a `ping` request with `pong`. The project includes the following components:  
1. A Python program for the ping-pong service.  
2. A Dockerfile to build a Docker image for the service.  
3. A `ping-pong-deployment.yaml` file to deploy the service in a Kubernetes cluster.  
4. A Makefile to automate the build, tag, push, and deploy processes.

---------------------
# Ping Pong Service

This service implements a simple Python script that responds to the `/ping` endpoint with the message `"pong"`. The service listens on port `5000`.

## Folder Structure

The code for this service is located in the `ping-pong-python` folder. It contains the following files:

### 1. `ping-pong-python.py`
This file contains the Python code to:
- Listen on port `5000`.
- Respond to the `/ping` endpoint with the message `"pong"`.

### 2. `Dockerfile`
This file contains the configuration to create a Docker image for the `ping-pong-python.py` script.

### 3. `ping-pong-deployment.yaml`
This file contains the Kubernetes deployment configuration for the Ping Pong service.

#### Breakdown of `ping-pong-deployment.yaml`:
- **apiVersion**: `apps/v1` (a stable API version for Kubernetes).
- **kind**: `Deployment` (to manage multiple replicas of pods running the Ping Pong service).
- **metadata**:
  - `name`: Name of the Deployment (`ping-pong-deployment`).
- **spec**:
  - **selector**:
    - `matchLabels`: Identifies the pod with the label `name: ping-pong-app`.
  - **template**:
    - **metadata**:
      - `labels`: Sets the pod label to `name: ping-pong-app`.
    - **spec**:
      - **containers**:
        - `image`: The container image (`ashwin0711/ping-pong-app`).
        - `name`: The container name (`ping-pong-app-container`).
        - `ports`:
          - `containerPort`: Exposes port `5000` from the container.

### 4. `Makefile`
This file automates the following tasks:
- Building the Docker image for the Ping Pong service.
- Deploying the containerized application to Kubernetes.

## How to Use
1. Clone the repository and navigate to the `ping-pong-python` folder.
2. Build the Docker image using the `Makefile` or manually with the `Dockerfile`.
3. Deploy the service to Kubernetes using the `ping-pong-deployment.yaml` file.

---
Feel free to contribute or raise issues in the repository!


