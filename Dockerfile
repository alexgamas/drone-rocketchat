FROM plugins/base:multiarch

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Rocket.Chat" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

COPY drone-rocketchat /bin/drone-rocketchat

ENTRYPOINT ["/bin/drone-rocketchat"]
