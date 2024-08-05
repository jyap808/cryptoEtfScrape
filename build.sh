#!/bin/bash

VERSION=1.0
APP='cryptoetfscrape'

docker build -t julian/${APP}:${VERSION} .
