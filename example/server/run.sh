#!/bin/bash

echo "Environment: $APP_ENV";

if [ "$APP_ENV" = "production" ]; \
        then \
        ./goproxy; \
        else \
        go get && \
        go get github.com/cespare/reflex && \
        reflex -r '\.go|json$' -s -- sh -c 'go build -o goproxy && DEBUG=true ./goproxy'; \
fi