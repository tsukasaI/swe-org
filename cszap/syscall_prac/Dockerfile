FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    gcc \
    libc6-dev \
    gdb \
    make \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace
