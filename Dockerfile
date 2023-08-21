FROM golang:1.17 AS core

LABEL Project="totemfi-tokenbridge-offchain"

# Set up apk dependencies
RUN apt update &&\
    apt install -y tzdata

ARG REPO_TOKEN

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute
## any further commands inside our /app
## directory
WORKDIR /app
## Add this go mod download command to pull in any dependencies
# RUN go mod download
RUN git config --global --add url."https://$REPO_TOKEN@github.com".insteadOf "https://github.com"
# RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig

RUN go mod vendor

RUN export CGO_CPPFLAGS="-I /usr/local/include"

ENV TZ=UTC

# Run the app
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags $CGO_CPPFLAGS -a -installsuffix cgo -o bridge-offchain

