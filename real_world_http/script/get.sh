#!/bin/sh

curl http://localhost:18888

curl -G --data-urlencode "query=hello world" http://localhost:18888

curl --head http://localhost:18888