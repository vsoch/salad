FROM golang:1.13-alpine3.10 as builder

#
# alpine image with the go tools
#

WORKDIR /code
ENV SRC_DIR=/go/src/github.com/vsoch/salad/
ADD . $SRC_DIR
WORKDIR $SRC_DIR

# Dependencies
RUN go mod vendor
RUN go build -o salad && cp salad /code/

FROM alpine:3.10
LABEL Maintainer vsochat@stanford.edu
COPY --from=builder /code/salad /code/salad

ENTRYPOINT ["/code/salad"]
CMD ["fork"]
