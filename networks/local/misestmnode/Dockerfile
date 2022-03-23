FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/mises-id/mises-tm

# Add source files
COPY . .

# build Ethermint
RUN make build-misestm-linux

# Final image
FROM golang:1.16 as final

WORKDIR /

RUN  sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN echo "deb http://mirrors.aliyun.com/debian/ buster main non-free contrib" > /etc/apt/sources.list

RUN  apt-get clean

RUN apt-get update

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/mises-id/mises-tm/build/misestmd /usr/bin/misestmd
COPY --from=build-env /go/src/github.com/mises-id/mises-tm/scripts/start.sh /

EXPOSE 26656 26657 1317 8545 8546

# Run ethermintd by default, omit entrypoint to ease using container with ethermintcli
ENTRYPOINT ["/bin/bash", "-c"]