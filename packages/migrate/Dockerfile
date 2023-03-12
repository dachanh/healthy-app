
FROM golang:1.20-alpine

RUN apk add -q --update \
    && apk add -q \
            bash \
            git \
            curl \ 
            make \
    && rm -rf /var/cache/apk/*

ENV ROOT /packages/migrate

WORKDIR $ROOT/postgres
COPY go.mod go.sum tool.go ./

WORKDIR $ROOT

COPY Makefile ./

COPY run.sh /

RUN  chmod 755 /run.sh

CMD ["/run.sh"]
