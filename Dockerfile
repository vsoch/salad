FROM iron/go:dev as builder

#
# alpine image with the go tools
#

WORKDIR /code
ENV SRC_DIR=/go/src/github.com/vsoch/salad/
ADD . $SRC_DIR
WORKDIR $SRC_DIR

# Dependencies
RUN go get github.com/urfave/cli
RUN go build -o salad && cp salad /code/

FROM alpine:3.7
LABEL Maintainer vsochat@stanford.edu
COPY --from=builder /code/salad /code/salad

ENTRYPOINT ["/code/salad"]
CMD ["fork"]
