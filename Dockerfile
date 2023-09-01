FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o our-diary our-diary.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/our-diary /build/our-diary

CMD [". /our-diary"]