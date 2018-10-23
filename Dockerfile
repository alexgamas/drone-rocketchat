FROM alpine:3.2
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD drone-rocketchat-plugin /bin/
ENTRYPOINT ["/bin/rocketchat-plugin"]
