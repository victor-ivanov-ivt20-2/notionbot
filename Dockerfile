FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o our-diary cmd/ourdiary/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/our-diary /build/our-diary

CMD [". /our-diary"]
