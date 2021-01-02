# We use Go 1.15. Do not upgrade this. Updates should be done by maintainers.
ARG GO_VERSION=1.15

# Alpine
FROM qmcgaw/godevcontainer:alpine

# Use root temporarily
USER root

# Install Alpine packages
RUN apk add --update \
  make \
  gnupg \
  nodejs \
  npm

# Install NPM packages
RUN npm i -g \
  yarn

# Install Go packages
RUN go get -v \
  github.com/uudashr/gopkgs/v2/cmd/gopkgs

# Embedded SQLite3 and File-based SQLite3 require CGO. :(
ENV CGO_ENABLED=1
