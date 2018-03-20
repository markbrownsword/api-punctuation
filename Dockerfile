FROM golang:1.10 AS build-env

# Setup working Directory
COPY . "/go/src/github.com/markbrownsword/faraway-api-punctuation"
WORKDIR "/go/src/github.com/markbrownsword/faraway-api-punctuation"

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure -vendor-only

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o goapp

#############
# Final image
FROM alpine

WORKDIR /app

COPY --from=build-env /go/src/github.com/markbrownsword/faraway-api-punctuation/goapp /app/

RUN apk --no-cache add ca-certificates

CMD ./goapp
