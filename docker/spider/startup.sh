#!/bin/sh
rc-service root start
rc-service atd restart
echo "cd /go/src/spider && go run main.go -mode 2 -p 1" | at now + 1minutes
crond -f