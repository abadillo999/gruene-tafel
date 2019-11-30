FROM alpine:edge

RUN apk add --update --no-cache \
    python3 \
    python3-dev \
    py-pip \
    build-base
RUN pip3 install \
    python-http-client \
    pylint \
    pytest \
    pytest-cov
RUN pip3 install virtualenv