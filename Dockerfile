ARG build_image=golang:1.13-alpine
ARG base_image=alpine

# -----------------
# # BINARY BUILD STAGE
# -----------------
FROM ${build_image} as build_stage

ENV GO111MODULE=on

RUN go get -v github.com/kolesa-team/http-api-mock

ENTRYPOINT ["/go/bin/http-api-mock","-config-path","/config"]

# -----------------
# IMAGE BUILD STAGE
# -----------------
FROM ${base_image} as final_stage

# Copy binary from build stage
COPY --from=build_stage /go/bin/http-api-mock /usr/local/bin/

RUN set -xe \
 && mkdir -p config data \
 && chown daemon:daemon config data

VOLUME /config
VOLUME /data
EXPOSE 8080 8081
USER daemon

ENTRYPOINT /usr/local/bin/http-api-mock -config-path /config -config-persist-path /data -server-port 8080 -console-port 8081
