# Project Overview

This project demonstrates the basic workings of Docker and Kubernetes, focusing on:  
1. A simple Hello-world Docker container.  
2. A Python Flask ping-pong service that responds to "ping".
3. A Go program for pattern detection in string.
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
### Deploying with Kubernetes
1. **Apply the Kubernetes Configuration**:
  The hello-world.yaml file contains deployment specifications. Used it to deploy the application to a Kubernetes cluster:
```bash
kubectl apply -f hello-world-docker/hello-world.yaml
```

-----------------

## Ping Pong Service

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

# Go Program: Pattern Detection

## Overview
This program checks if a string follows a specific pattern:
- Starts with `"i"` or `"I"`.
- Contains `"a"` or `"A"` anywhere in the middle.
- Ends with `"n"` or `"N"`.

## Code Location
The code for this program is located in the `main.go` file.

## How It Works
1. Reads an input string.
2. Evaluates the string against the defined pattern.
3. Returns whether the string matches the pattern.

