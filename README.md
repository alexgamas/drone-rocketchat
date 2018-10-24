![rocket.chat-logo](logo.svg)


# drone-rocketchat

[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Report](https://goreportcard.com/badge/github.com/alexgamas/drone-rocketchat)](https://goreportcard.com/report/github.com/alexgamas/drone-rocketchat)
[![](https://images.microbadger.com/badges/image/alexgamas/drone-rocketchat.svg)](https://microbadger.com/images/alexgamas/drone-rocketchat "Get your own image badge on microbadger.com")

Drone plugin for sending Rocket.Chat notifications. For the usage information and a listing of the available options please take a look at [the docs](https://github.com/alexgamas/drone-rocketchat/blob/release/0.1/DOCS.md).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-rocketchat
docker build --rm -t plugins/rocketchat .
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_URL=https://my.rocketchat.com/ \
  -e PLUGIN_CHANNEL=ci \
  -e PLUGIN_USERNAME=ci \
  -e DRONE_REPO_OWNER=ci \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=ci \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=https://github.com/alexgamas/drone-rocketchat \
  -e DRONE_TAG=0.0.1 \
  plugins/rocketchat
```
