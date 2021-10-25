#!/bin/bash
kill -9 `lsof -ti:8080`
nohup go run ../main.go &