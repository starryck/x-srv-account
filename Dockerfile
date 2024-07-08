FROM golang:1.22.5-alpine3.20 as builder
LABEL maintainer=gordon.lai@starryck.com

RUN apk update && \
    apk add --no-cache git openssh gcc musl-dev && \
    git config --global url."git@github.com:".insteadOf "https://github.com/"

COPY . /source
WORKDIR /source
RUN go build -o ./main.exe -tags=musl,sonic,avx .

FROM alpine:3.20

ARG git_tag
ARG git_commit
ENV GIT_TAG=${git_tag}
ENV GIT_COMMIT=${git_commit}

COPY --from=builder /source/main.exe .
RUN ln -s /main.exe /usr/local/bin/service

CMD ["service"]
