

build:
	docker build -t ping-pong-app .

tag:
	docker tag ping-pong-app ashwin0711/ping-pong-app:latest

push:
	docker push ashwin0711/ping-pong-app:latest

run:
	docker run -p 5000:5000 ping-pong-app

deploy:
	kubectl apply -f ping-pong-deployment.yaml


all: build tag push run deploy 
