FROM golang:1.23.2-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# dependencies
COPY ["app/go.mod","app/go.sum", "./"]
RUN go mod download

# build
COPY app ./
RUN go build -o ./bin/app cmd/app/main.go 

FROM alpine

COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.yaml /config.yaml
# COPY configs/.env /.env

CMD [ "/app" ]