# drone-rocketchat

[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-rocketchat)](https://goreportcard.com/report/github.com/drone-plugins/drone-rocketchat)
[![](https://images.microbadger.com/badges/image/plugins/rocketchat.svg)](https://microbadger.com/images/plugins/rocketchat "Get your own image badge on microbadger.com")

Drone plugin for sending Rocket.Chat notifications. For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-rocketchat/).

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

### Parameters

Plugin parameters:


### Secrets


### Logos



### Docs


### Images


### Testing


### Vendoring

Please vendor dependencies in a manner compatible with `GOVENDOREXPERIMENT`. All official drone plugins should use [govend](https://github.com/govend/govend) with the `--prune` flag.

The Rocket.Chat plugin posts build status messages to your channel. The below pipeline configuration demonstrates simple usage:

```yaml
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
```

Example configuration with custom username:

```diff
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
+   username: drone
```

Example configuration with custom avatar:

```diff
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
+   icon_url: https://unsplash.it/256/256/?random
```

Example configuration with image attachment:

```diff
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
+   image_url: https://unsplash.it/256/256/?random
```

Example configuration for success and failure messages:

```diff
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
+   when:
+     status: [ success, failure ]
```

Example configuration with a custom message template:

```diff
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
+   template: >
+     {{#success build.status}}
+       build {{build.number}} succeeded. Good job.
+     {{else}}
+       build {{build.number}} failed. Fix me please.
+     {{/success}}
```

Example configuration with a custom message template linking usernames and channels:

```diff
pipeline:
  rocketchat:
    image: plugins/rocketchat
    url: https://my.rocketchat.com/
    channel: ci
+   template: >
+     {{#success build.status}}
+       build {{build.number}} succeeded. Good job. <@john.doe>
+     {{else}}
+       build {{build.number}} failed. Fix me please. <@channelname> <@someone>
+     {{/success}}
```

# Parameter Reference

url
: incoming [url](https://my.rocketchat.com/) for posting to a channel

channel
: messages sent to the above url are posted here

username
: choose the username this integration will post as

template
: overwrite the default message template

image_url
: a valid URL to an image file that will be displayed inside a message attachment

icon_url
: a valid URL that displays a image to the left of the username

icon_emoji
: displays a emoji to the left of the username

# Template Reference

repo.owner
: repository owner

repo.name
: repository name

build.status
: build status type enumeration, either `success` or `failure`

build.event
: build event type enumeration, one of `push`, `pull_request`, `tag`, `deployment`

build.number
: build number

build.commit
: git sha for current commit

build.branch
: git branch for current commit

build.tag
: git tag for current commit

build.ref
: git ref for current commit

build.author
: git author for current commit

build.link
: link the the build results in drone

build.created
: unix timestamp for build creation

build.started
: unix timestamp for build started

build.pull
: pull request number (empty string if not a pull request)

build.deployTo
: env that the build was deployed to.

# Template Function Reference

uppercasefirst
: converts the first letter of a string to uppercase

uppercase
: converts a string to uppercase

lowercase
: converts a string to lowercase. Example `{{lowercase build.author}}`

datetime
: converts a unix timestamp to a date time string. Example `{{datetime build.started}}`

success
: returns true if the build is successful

failure
: returns true if the build is failed

truncate
: returns a truncated string to n characters. Example `{{truncate build.sha 8}}`

urlencode
: returns a url encoded string

since
: returns a duration string between now and the given timestamp. Example `{{since build.started}}`
