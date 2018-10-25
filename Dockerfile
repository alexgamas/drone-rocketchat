FROM plugins/base:multiarch

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Rocket.Chat" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

COPY release/linux/amd64/drone-rocketchat /bin/

ENTRYPOINT ["/bin/drone-rocketchat"]
