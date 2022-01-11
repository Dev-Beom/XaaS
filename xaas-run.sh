#!/bin/sh

cd APIServer || exit
./api-server-run.sh &

cd ../ControlManager || exit
go run main.go