ARG baseImage="ubuntu:bionic"
FROM ${baseImage} as base

# Add the magma apt repo
RUN apt-get update && \
    apt-get install -y apt-utils software-properties-common apt-transport-https
COPY orc8r/tools/ansible/roles/pkgrepo/files/jfrog.pub /tmp/jfrog.pub
RUN apt-key add /tmp/jfrog.pub && \
    apt-add-repository "deb https://facebookconnectivity.jfrog.io/artifactory/list/dev/ bionic main"

# Install the runtime deps.
RUN apt-get update && apt-get install -y \
    bzr \
    curl \
    daemontools \
    gcc \
    git \
    libc-ares-dev \
    libev-dev \
    libevent-dev \
    libffi-dev \
    libjansson-dev \
    libjemalloc-dev \
    libssl-dev \
    libsystemd-dev \
    magma-nghttpx=1.31.1-1 \
    make \
    net-tools \
    pkg-config \
    python-cffi \
    python3-pip \
    redis-server \
    rsyslog \
    sudo \
    unzip \
    vim \
    virtualenv

# Install Golang 1.12.
WORKDIR /usr/local
RUN curl https://facebookconnectivity.jfrog.io/artifactory/generic/go1.12.4.linux-amd64.tar.gz -O && \
    tar xf go1.12.4.linux-amd64.tar.gz && \
    cp -r go/bin/* /usr/local/bin/

# Install protobuf compiler.
RUN curl -Lfs https://github.com/google/protobuf/releases/download/v3.1.0/protoc-3.1.0-linux-x86_64.zip -o protoc3.zip && \
    unzip protoc3.zip -d protoc3 && \
    mv protoc3/bin/protoc /bin/protoc && \
    chmod a+rx /bin/protoc && \
    mv protoc3/include/google /usr/include/ && \
    chmod -R a+Xr /usr/include/google && \
    rm -rf protoc3.zip protoc3

ENV GOBIN /var/opt/magma/bin
ENV MAGMA_ROOT /magma
ENV PIP_CACHE_HOME ~/.pipcache
ENV PYTHON_BUILD /build/python
ENV PATH ${PYTHON_BUILD}/bin:${PATH}:${GOBIN}
ENV GO111MODULE on
# Use public go modules proxy
ENV GOPROXY https://proxy.golang.org

RUN printenv > /etc/environment


# Copy just the go.mod and go.sum files to download the golang deps.
# This step allows us to cache the downloads, and prevents reaching out to
# the internet unless any of the go.mod or go.sum files are changed.
COPY modules/cwf/cloud/go/go.* $MAGMA_ROOT/modules/cwf/cloud/go/
COPY modules/cwf/gateway/go.* $MAGMA_ROOT/modules/cwf/gateway/
COPY modules/lte/cloud/go/go.* $MAGMA_ROOT/modules/lte/cloud/go/
COPY modules/feg/cloud/go/go.* $MAGMA_ROOT/modules/feg/cloud/go/
COPY modules/feg/cloud/go/protos/go.* $MAGMA_ROOT/modules/feg/cloud/go/protos/
COPY modules/feg/gateway/go.* $MAGMA_ROOT/modules/feg/gateway/
COPY lib/go/go.* $MAGMA_ROOT/lib/go/
COPY lib/go/protos/go.* $MAGMA_ROOT/lib/go/protos/
COPY orc8r/cloud/go/go.* $MAGMA_ROOT/orc8r/cloud/go/
COPY gateway/go/go.* $MAGMA_ROOT/gateway/go/
WORKDIR $MAGMA_ROOT/cwf/gateway
RUN go mod download; exit 0
# Install protoc-gen-go
RUN go install github.com/golang/protobuf/protoc-gen-go; exit 0

# Symlink python scripts.
RUN ln -s /build/python/bin/generate_service_config.py /usr/local/bin/generate_service_config.py
RUN ln -s /build/python/bin/generate_nghttpx_config.py /usr/local/bin/generate_nghttpx_config.py

# Build the code.
COPY modules/cwf $MAGMA_ROOT/modules/cwf
COPY modules/feg $MAGMA_ROOT/modules/feg
COPY modules/lte/cloud $MAGMA_ROOT/modules/lte/cloud
COPY lib/go $MAGMA_ROOT/lib/go
COPY orc8r/cloud $MAGMA_ROOT/orc8r/cloud
COPY gateway/go $MAGMA_ROOT/gateway/go

# -----------------------------------------------------------------------------
# Builder image with binary
# -----------------------------------------------------------------------------
FROM base as builder
# Enable make gen if proto gen is required
# RUN make -C $MAGMA_ROOT/modules/cwf/gateway gen
RUN make -C $MAGMA_ROOT/modules/cwf/gateway build

# -----------------------------------------------------------------------------
# Production image
# -----------------------------------------------------------------------------
FROM ubuntu:bionic AS cwag_go

# Install envdir.
RUN apt-get -y update && apt-get -y install daemontools curl iptables

# Copy the build artifacts.
COPY --from=builder /var/opt/magma/bin /var/opt/magma/bin

# Copy the configs.
COPY cwf/gateway/configs /etc/magma

# Create empty envdir directory
RUN mkdir -p /var/opt/magma/envdir

RUN mkdir -p /var/opt/magma/configs
