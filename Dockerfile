FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./cmd/api/main.go

EXPOSE 3003

CMD [ "/build" ]