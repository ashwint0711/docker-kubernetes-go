#Python runtine
FROM python:3.8-slim

workdir /app

add . /app

run pip install flask

expose 5000

cmd ["python3", "./ping-pong-app.py"]
