#!/bin/sh

cd APIServer
./api-server-run.sh &

cd ../ControlManager
go run main.go