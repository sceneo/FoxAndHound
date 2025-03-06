#!/bin/sh

cd database
docker-compose up -d &
cd ../backend
go run main.go &
cd ../frontend
npm start
