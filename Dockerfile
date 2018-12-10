FROM iron/go:dev

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
ENTRYPOINT ["/code/salad"]
CMD ["fork"]
