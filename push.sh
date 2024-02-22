#!/bin/bash
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

docker buildx build --platform linux/arm64,linux/amd64 -t "$IMAGEURL" -f cmd/scrap/Dockerfile --push . 

# docker push "$IMAGEURL"
