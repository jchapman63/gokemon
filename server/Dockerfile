# server is to be built from the client but with the detached mode option

FROM golang:latest

WORKDIR /app

# fist install source code dependencies
COPY go.mod go.sum ./
RUN go mod download

# then install the source code
COPY *.go ./

# create an application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /gokemon

# TODO: Go app pulls this port in with OS library?
EXPOSE 8081

CMD ["/gokemon"]

# TODO configure main of gokemon to run server