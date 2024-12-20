#!/bin/bash

VERSION=1.0
APP='cryptoetfscrape'

docker run -dit \
	--restart always \
        --env-file env \
        --name ${APP}-${VERSION} \
        julian/${APP}:${VERSION}
