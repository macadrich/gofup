# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# create directory
RUN mkdir /go/src/gofup

# Copy the local package files to the container's workspace.
ADD . /go/src/gofup

ADD . /go

# set workspace
WORKDIR /go/src/gofup

# Build command inside the container.
RUN go install /go/src/gofup/cli

ENTRYPOINT ["/go/bin/cli"]
#CMD [ "/go/bin/cli" ]

EXPOSE 15223

# VOLUME [ "/go/src/gofup/myfile" ]
# enable to access container
