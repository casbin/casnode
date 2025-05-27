FROM golang:1.21.13 AS BACK
WORKDIR /go/src/casnode
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server . \
    && apt update && apt install wait-for-it && chmod +x /usr/bin/wait-for-it

FROM node:18.19.0 AS FRONT
WORKDIR /web
COPY ./web .
RUN yarn install --frozen-lockfile --network-timeout 1000000 && NODE_OPTIONS="--max-old-space-size=4096" yarn run build

FROM alpine:latest
RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
LABEL MAINTAINER="https://casnode.org/"

COPY --from=BACK /go/src/casnode/ ./
COPY --from=BACK /usr/bin/wait-for-it ./
RUN mkdir -p web/build && apk add --no-cache bash coreutils
COPY --from=FRONT /web/build /web/build
CMD ./wait-for-it db:3306 -- ./server