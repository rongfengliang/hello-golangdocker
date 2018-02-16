# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/rongfengliang/fndemo
ADD . $D
RUN cd $D && go build -o fn-hello && cp fn-hello /tmp/

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /tmp/fn-hello /app/hello
CMD ["./hello","demoapp"]