@echo off
docker build --build-arg HTTP_PROXY=http://127.0.0.1:7890  -t monitor-app:0.0.1 . 